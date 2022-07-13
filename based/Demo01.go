package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name   string "user name" //这引号里面的就是tag
	Passwd string "user passsword"
}

type UserArray []*User

func main() {
	const Name = "Ecomb"
	var count = 1
	var flag = "go"
	fmt.Println("hello", count, flag)

	var arr [5]int

	for i := 0; i < 2; i++ {
		arr[i] = i
	}
	i := 0
	i++
	fmt.Println(arr[i])

	for i := 0; i < 2; i++ {
		fmt.Println(arr[i])
	}

	user := &User{"chronos", "pass"}
	s := reflect.TypeOf(user).Elem() //通过反射获取type定义

	fmt.Println("String", s.String())

	var num = s.NumField()
	var headers = make([]string, num)

	for i := 0; i < num; i++ {
		fmt.Println(s.Field(i).Tag)
		headers[i] = string(s.Field(i).Tag)
	}

	user01 := &User{
		Name: "zhou",
		Passwd: "fa",
	}

	value := reflect.ValueOf(*user)
	cou := value.NumField()
	for i := 0; i < cou; i++ {
		f := value.Field(i)
		switch f.Kind() {
		case reflect.String:
			fmt.Println("1:", f.String())
		case reflect.Int:
			fmt.Println(f.Int())
		}
	}

	user02 := &User{
		Name: "xiao",
		Passwd: "qwqq",
	}

	var userArray UserArray
	userArray = append(userArray, user01)
	userArray = append(userArray, user02)

	for index, user := range userArray {
		u := reflect.TypeOf(user).Elem()
		//value := reflect.ValueOf(user)

		fmt.Println(index, ":", user, "-> ", u.Field(0))
	}

	//rows := make([][]string, len(userArray), len(getHeaders()))
	var rows [][]string
	for _, csv := range userArray {
		value := reflect.ValueOf(*csv)
		cou := value.NumField()
		row := make([]string, cou)
		for i := 0; i < cou; i++ {
			f := value.Field(i)
			if f.Kind() == reflect.String {
				row[i] = f.String()
			}else {
				row[i] = ""
			}
		}
		rows = append(rows, row)
	}
	fmt.Println(len(rows), ":", rows[0][0], ", ", len(getHeaders()))
}


func getHeaders() []string {
	csvReflectType := reflect.ValueOf(User{}).Type()
	var num = csvReflectType.NumField()
	var headers = make([]string, num)

	for i := 0; i < num; i++ {
		fmt.Println(csvReflectType.Field(i).Tag)
		headers[i] = string(csvReflectType.Field(i).Tag)
	}
	return headers
}