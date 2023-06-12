package main
import "fmt"

func main()  {
	fmt.Print("my weight on the surface of Mars is ")
	fmt.Print(180 * 0.37)
	fmt.Println("my weight is :" , 180 * 0.37)

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
} 