package main

import (
	"fmt"
	"reflect"
)

type User struct {
	name string
	age int
}
type Admin struct {
	Name     string "user name" //这引号里面的就是tag
	PassWord string "user password"
}

func main() {
	user := User{"zhouzg", 2}
	fmt.Println(user)

	admin := &Admin{"chronos", "pass"}
	s := reflect.TypeOf(admin).Elem() //通过反射获取type定义
	for i := 0; i < s.NumField(); i++ {
		fmt.Println(s.Field(i).Tag) //将tag输出出来
	}
}