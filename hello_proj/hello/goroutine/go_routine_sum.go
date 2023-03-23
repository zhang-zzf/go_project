package main

import "fmt"

func sum(arr []int, ch chan int) {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	ch <- sum
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	ch := make(chan int)
	go sum(arr[:len(arr)/2], ch)
	go sum(arr[len(arr)/2:], ch)
	s1, s2 := <-ch, <-ch
	fmt.Printf("%v %v = %v+%v\n", arr, s1+s2, s1, s2)
}
