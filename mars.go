package main
import (
	"fmt"
	"math/rand"
)

func main()  {
	fmt.Print("my weight on the surface of Mars is ")
	fmt.Print(180 * 0.37)
	fmt.Println("my weight is :" , 200 * 0.37)

	fmt.Printf("%-15v $%4v\n","spaceX",95)
	fmt.Printf("%-15v $%4v\n","verygood",100)

	var (
		num = 1
		age = 10
	)
	const time,name = 120,"lzl"

	fmt.Println(num,age)
	fmt.Println(time,name)

	// time = 10
	num += 1

	fmt.Print(time , num) 

	const l = 56000000
	fmt.Println("飞船最低速度为：", l/(28*24) , "km/h")

	var number = rand.Intn(10) + 1
	fmt.Print(number)
} 