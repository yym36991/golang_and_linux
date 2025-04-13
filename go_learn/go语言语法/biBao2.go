package main

import (
	"fmt"
)

type Student struct {
	Name string
	Age  int
}

func InitPools(configs []Student) []Student {
	var students []Student

	for _, config := range configs {
		dailFunc := func() func() Student {
			return func() Student {
				// 注释掉 configCopy 来展示闭包捕获引用的问题
				return Student{
					Name: config.Name,
					Age:  config.Age * 2,
				}
			}
		}()
		students = append(students, dailFunc())
	}

	return students
}

func main() {
	configs := []Student{
		{"Tom", 20},
		{"Jerry", 30},
		{"Lily", 40},
	}

	pools := InitPools(configs)

	for i, pool := range pools {
		fmt.Printf("Pool %d: %v\n", i, pool)
	}
}
