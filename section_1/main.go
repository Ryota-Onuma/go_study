package main

import (
	"fmt"
	"math/rand"
	"math"
)
var python, java =  true, false
var f,g int = 1,2
func add(x int, y int) int {
	return x  +  y
}

func swap(x,y string) (string,string){
	return y , x
}

func split(sum int) (x,y int){
	x = sum * 4 /9
	y = sum - x
	return
}

func main() {
	fmt.Println("Hello, 世界")
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Println("円周率は", math.Pi)
	num_1 := 5
	num_2 := 50
	fmt.Println(add(num_1,num_2))
	a, b := swap("hello","world")
	fmt.Println(a,b)
	c, d := split(100)
	fmt.Println(c,d)

	var i int
	fmt.Println(i,c,python,java)

	const TheWorld = "ザ ワールド"
	fmt.Println(TheWorld)
}

