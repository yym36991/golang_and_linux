package main

import (
	"fmt"
)

func main() {
	sum := Sum(1, 2)
	fmt.Println(sum(1, 2))
	fmt.Println(sum(3, 4))
	fmt.Println(sum(5, 6))
}

func Sum(a, b int) func(int, int) int {
	return func(a, b int) int {
		return a + b
	}
}

/*
1、执行到main函数内的sum变量，Sum(1, 2) 时，Sum产生了自己的funcval结构体指针val1,保存自己的入口地址和环境信息。
接着：在执行到匿名函数时，匿名函数会有一个自己的funcval结构体val2，在返回时，val2保存了匿名函数的入口地址，和捕获的自由变量a,b(合称闭包)。
当Sum返回时，val1被回收。返回的val2（闭包）被赋值给sum变量。
 匿名函数和捕获的自由变量a,b，一起叫做闭包。


2、闭包的定义：指一个函数及其捕获的外部变量的组合。

3、所有只有匿名函数并不构成闭包。匿名函数捕获它外部的自由变量之后，才变身闭包。

*/

type funcval struct {
	fn    uintptr //指向函数入口的指针
	extra *uint8  //指向存储数据的指针
}

func (s *deleteServer) InitRedisPools(conf *config.Config) error {
	redisConfigs := conf.DeleteServer.RedisConfigs
	for _, redisConf := range redisConfigs {
		dailFunc := func() func() (redis.Conn, error) {
			return func() (redis.Conn, error) {
				conn, err := redis.Dial("tcp", redisConf.Address,
					redis.DialClientName(conf.DeleteServer.RedisUser),
					redis.DialPassword(redisConf.Password))
				if err != nil {
					return nil, err
				}
				return conn, nil
			}
		}()

		redisPool := &redis.Pool{Dial: dailFunc, MaxIdle: 10}

		conn := redisPool.Get()
		if _, err := conn.Do("PING"); err != nil {
			return fmt.Errorf("cannot connect to redis %s: %v", redisConf.Address, err)
		}
		conn.Close()
		s.redisPools = append(s.redisPools, redisPool)
	}
	return nil
}
