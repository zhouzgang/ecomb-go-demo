package main

import "fmt"

func main() {
	const Name = "Ecomb"
	var count int = 1
	var flag = "go"
	fmt.Println("hello", count, flag)

	var arr [5]int

	for i := 0; i < 2; i++ {
		arr[i] = i
	}

	for i := 0; i < 2; i++ {
		fmt.Println(arr[i])
	}
}