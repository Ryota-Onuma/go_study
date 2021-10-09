package main

import "fmt"

type Vertex struct {
	X int
	Y int
}



func main(){
	i := 42
	p := &i
	fmt.Println(i)
	fmt.Println(p)
	fmt.Println(*p)
	*p = 1000
	fmt.Println(i) // i の値が変わっている

	v := 	Vertex{100,200}
	pv := &v
	fmt.Println(v)
	v.Y = 1
	pv.X = 1e9
	fmt.Println(v)
	fmt.Println(*pv)
}
