package main

import "fmt"
import "time"

func fizzBuzz(){
	sum := 1
	for sum < 1000 {
		if sum % 15 == 0 {
			fmt.Println("FizzBuzz")
		} else if sum % 3 == 0 {
			fmt.Println("Buzz")
		}else if sum  % 5 ==  0{
			fmt.Println("Fizz")
		}
		sum++
	}
}

func judgeDate(){
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday  {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tommorow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("It is too far away")
	}
}

func main() {
	defer fmt.Println("finished")
	var sum int = 0
	for i :=0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
	sum = 1
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)
	sum = 1
	for sum < 10000{
		sum += sum
	}
	fmt.Println(sum)
	fizzBuzz()
	judgeDate()

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	
}
