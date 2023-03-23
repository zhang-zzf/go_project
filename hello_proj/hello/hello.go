package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	reversed_str := stringutil.Reverse("Hello")
	fmt.Println(reversed_str)
	upper_str := stringutil.ToUpper("Hello, World.")
	fmt.Println(upper_str)
}
