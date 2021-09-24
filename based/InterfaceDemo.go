package main

// 参考链接：https://sanyuesha.com/2017/07/22/how-to-understand-go-interface/

import "fmt"

//1
type I interface {
	Get() int
	Set(int)
}

//2
type S struct {
	Age int
}

func(s S) Get() int {
	fmt.Println(s)
	return s.Age
}

// todo 这里的 s *S 怎么理解
func(s *S) Set(age int) {
	fmt.Println(s)
	s.Age = age
}

//3
func f(i I){
	i.Set(10)
	fmt.Println(i.Get())
}

func main() {
	s := S{}
	f(&s)  //4
}