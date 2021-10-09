package main
import (
	"fmt"
	"strings"
	"strconv"
)

func practiceArray(){
	var array[2]string
	array[0] = "Hello"
	array[1] = "World"
	fmt.Println(array[0],array[1])
	numbers := [6]int{1,2,3,4,5,6}
	fmt.Println(numbers)
	var slice []int = numbers[0:4]
	fmt.Println(slice)

	users  := []struct{
		Name string
		Tel string
		Age int}{
		{"hoge太郎","0123456789",24},
		{"hoge二郎","0123456783",23},
		{"hoge三郎","0123456781",21},
		{"hoge四郎","0123456781",21},
	}

	fmt.Printf("len=%d cap=%d %v\n", len(users), cap(users), users)
	fmt.Println(users)

	animals := make([]string,3) 

	fmt.Println(animals)
 
	for i := 0; i < 3; i++ {
		var str_build strings.Builder
		str_build.WriteString("クマ")
		str_build.WriteString(strconv.Itoa(i))
		animals[i] = str_build.String()
	}
	fmt.Println(cap(animals),animals)
	animals = append(animals, "クマ100")
	fmt.Println(cap(animals),animals)

	for i,v := range animals {
		fmt.Printf("indexは%v value=%v\n", i, v)
	}
}

func main(){
	practiceArray()
}
