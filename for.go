package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	var room= "lake"
	
	switch {
	case room == "cave":
		fmt.Println("人住在洞穴")
	case room == "lake":
		fmt.Println("睡湖里")
		fallthrough		//执行下一个case体里的内容
	case room == "dahaosi":
		fmt.Println("人住在大豪斯")
	}

	var count = 3

	for count > 0 {
		fmt.Println(count)
		// 每一秒停顿一下（秒数可自定）
		time.Sleep(time.Second)		//time.Sleep( n * time.Second )
		count --
	}
	fmt.Println("睡觉啦！")


	const a = 88
	//  获取1-100的整数
	var n = rand.Intn(100)
	for {
		if n > a {
			fmt.Println("大了")
			n = rand.Intn(100)
		}else if n < a{
			fmt.Println("小了")
			n = rand.Intn(100)
		}else{
			fmt.Println("ok")
			break
		}
	}
	

}