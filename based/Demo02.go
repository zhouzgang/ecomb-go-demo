package main

import (
	"fmt"
)


func main() {
	var titleList  = []string{"User Name", "Company Name", "Company ID", "Date/Time of Generation", "Date Range"}
	for key := range titleList {
		fmt.Println(titleList[key])

	}
}