package main

import (
	"fmt"
	"strings"
)

func main()  {
	fmt.Println("你发现自己在洞穴")

	var command = "走出去"
	// 判断里面是否含有 “走”
	var exit = strings.Contains(command,"走")

	fmt.Println("你离开了洞穴",exit)

	var age = 15
	var id = 996
	var game = ""
	var run = ""

	if age >= 18{
		game = "成年"
	}else{
		game = "未成年"
	}

	println("你是否成年？" ,"我：\"", game,"\" ")

	if age >= 18 || age/2 == 0 {
		run = "go run"
	}else{
		run = "just run"
	}
	println(run)

	if id < 1007 && id/2 == 0 {
		run = "躺平"
	}else{
		run = "福报"
	}
	println(id,"是",run)

	var year = 2023
	var leap = year%400 == 0 || (year%4==0 && year%100 != 0)

	if leap {
		fmt.Println("今年是闰年")
	}else{
		fmt.Println("今年不是润年")
	}
}


