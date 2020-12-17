package main

import (
	"fmt"
	"time"
)

func main() {

	//创建一个周期性的定时器
	ticker := time.NewTicker(3 * time.Second)
	fmt.Println("当前时间为:", time.Now())

	go func() {
		for {

			//从定时器中获取数据
			t := <-ticker.C
			fmt.Println("当前时间为:", t)

		}
	}()

	time.Sleep(time.Hour)
}
