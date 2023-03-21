package main

import "fmt"
import "github.com/greetings"

func main() {
	fmt.Println("Hello, World.")
	msg := greetings.Hello("zhang.zzf")
	fmt.Println(msg)
}
