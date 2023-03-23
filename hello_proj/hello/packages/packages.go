package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	fmt.Println("My favorite num is ", rand.Intn(10))
	abs_num := math.Abs(-11)
	fmt.Println(abs_num)
	fmt.Println(math.Pi)
	sum_n := sum(-100, 599)
	fmt.Println(sum_n)
	fmt.Println(swap_string("World.", "Hello,"))
	c := 'A'
	fmt.Printf("'A' Type: %T \n", c)
}

func sum(x, y int) int {
	return x + y
}

func swap_string(s1, s2 string) (string, string) {
	return s2, s1
}
