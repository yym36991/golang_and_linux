package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"sort"
	"strconv"
	"sync"
	"syscall"
	"time"
)

var meetMarker = false

func getFileTimes(filePath string) (time.Time, time.Time, error) {
	var stat syscall.Stat_t
	err := syscall.Stat(filePath, &stat)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}

	modifyTime := time.Unix(stat.Mtim.Sec, stat.Mtim.Nsec) // 最近修改时间
	changeTime := time.Unix(stat.Ctim.Sec, stat.Ctim.Nsec) // 最近状态变更时间

	return modifyTime, changeTime, nil
}

func readFile(filePath string, wg *sync.WaitGroup, errChan chan<- string, markerPath string, maxFile *string, doneChan <-chan bool) {
	defer wg.Done()

	select {
	case <-doneChan:
		return
	default:
	}

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		modifyTime, createTime, statErr := getFileTimes(filePath)
		if statErr == nil {
			errMsg := fmt.Sprintf("path=%s; modify_time=%s; create_time=%s\n",
				filePath, modifyTime.Format(time.RFC3339), createTime.Format(time.RFC3339))
			errChan <- errMsg
		} else {
			errChan <- filePath
		}
		log.Printf("Error reading file %s: %v\n", filePath, err)
		return
	}

	_ = data

	*maxFile = filePath
	// log.Printf("filePath=%s; cur maxFile=%s\n", filePath, *maxFile)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run script.go <directory> [<number_of_goroutines>] [marker=\"/path/name.txt\"]")
		return
	}

	logFile, err := os.OpenFile("error_log.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	sumLogFile, err := os.OpenFile("sum.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer sumLogFile.Close()

	directory := os.Args[1]
	numGoroutines := 1
	var markerPath string
	if len(os.Args) >= 3 {
		numGoroutines, err = strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatalf("Invalid number of goroutines: %v\n", err)
		}
	}
	if len(os.Args) == 4 {
		markerPath = os.Args[3]
	}
	// 监控 Ctrl+C 信号
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	var wg sync.WaitGroup
	errChan := make(chan string, 100)
	sem := make(chan struct{}, numGoroutines)

	go func() {
		for errPath := range errChan {
			writeErrorToFile(errPath)
		}
	}()

	var NewMaxFile string
	doneChan := make(chan bool)
	// doneChan2 := make(chan bool)

	go func() {
		<-signalChan
		log.Printf("Received shutdown signal.")
		close(doneChan)
		// close(doneChan2)
	}()

	fileCount := 0

	stack := []string{directory}

	for len(stack) > 0 {
		currentDir := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		files, err := ioutil.ReadDir(currentDir)
		if err != nil {
			log.Fatalf("Failed to read directory: %v\n", err)
		}

		sort.Slice(files, func(i, j int) bool {
			return files[i].Name() > files[j].Name()
		})

		for _, file := range files {
			// log.Printf("file before doneChan ")
			select {
			case <-doneChan:
				log.Printf("Shutdown signal received, exiting inner loop.")
				goto end
			default:
			}
			// log.Printf("to traverse file: %s\n", file.Name())
			filePath := filepath.Join(currentDir, file.Name())
			if file.IsDir() {
				stack = append(stack, filePath)
			} else {
				if markerPath == "" || meetMarker {
					select {
					case <-doneChan:
						break
					default:
					}

					wg.Add(1)
					sem <- struct{}{}
					go func(filePath string) {
						defer func() { <-sem }()
						readFile(filePath, &wg, errChan, markerPath, &NewMaxFile, doneChan)

						// 记录文件数
						sumLogEntry := fmt.Sprintf("fileCount=%d; path=%s\n", fileCount, filePath)
						sumLogFile.WriteString(sumLogEntry)

					}(filePath)

					fileCount++
				} else {
					if filePath == markerPath {
						meetMarker = true
					}
				}
			}
		}
		// log.Printf("NewMaxFile=%s\n", NewMaxFile)
	}

	log.Printf("NewMaxFile=%s\n", NewMaxFile)

end:
	log.Printf("run end, NewMaxFile=%s\n", NewMaxFile)
	wg.Wait()
	log.Printf("wg.Wait done")
	close(errChan)
	writeMarkerFile(NewMaxFile)

	// 输出总计数
	log.Printf("Total files scanned: %d\n", fileCount)

}

func writeMarkerFile(newmaxFile string) {
	if newmaxFile != "" {
		markerFile, err := os.OpenFile("marker.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Printf("Could not open marker file: %v\n", err)
			return
		}
		defer markerFile.Close()

		if _, err := markerFile.WriteString(newmaxFile + "\n"); err != nil {
			log.Printf("Could not write to marker file: %v\n", err)
		}
	} else {
		log.Printf("newmaxFile is empty, nothing to log.\n")
	}
}

func writeErrorToFile(filePath string) {
	log.Printf("begin to writeErrorToFile to file: %s\n", filePath)
	errorLogFile := "error_dir.txt"
	f, err := os.OpenFile(errorLogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("Could not open error directory file: %v\n", err)
		return
	}
	defer f.Close()

	if _, err := f.WriteString(filePath + "\n"); err != nil {
		log.Printf("Could not write to error directory file: %v\n", err)
	}
}
