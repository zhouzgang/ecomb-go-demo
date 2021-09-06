package main

import "fmt"

type User struct {
	name string
	age int
}

func main() {
	user := User{"zhouzg", 2}
	fmt.Println(user)

}