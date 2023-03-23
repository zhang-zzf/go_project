package main

import "fmt"

type Vertex struct {
	X int
	Y int
	z int
}

func main() {
	v := Vertex{1, 2, 3}
	fmt.Println(v)
	fmt.Println(v.X, v.Y, v.z)
	p := &v
	p.X = 11
	fmt.Println(p.X, p.Y, p.z)
}
