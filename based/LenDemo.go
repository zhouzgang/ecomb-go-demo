package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"strconv"
)

func main() {

	testMap := make(map[string]interface{})
	testMap["name"] = "typ111"
	testMap["age"] = 1231
	testMap["addr"] = "beijing111"

	var rows []string
	rows = append(rows, "a")
	rows = append(rows, "b")
	rows = append(rows, "c")


	axis := len(testMap) + 2 + len(rows) + 1
	fmt.Println("axis", "A" + strconv.Itoa(axis))


	f := excelize.NewFile()
	// Create a new sheet.
	index := f.NewSheet("Sheet1")
	f.SetActiveSheet(index)

	//axis = "A" + strconv.Itoa(len(testMap) + 2 + len(rows) + 1)
	err := f.SetCellValue("Sheet1", "A" + strconv.Itoa(axis), "Data in blue cells are premium features. It is currently made free to you for a limited time only. Lalamove reserves the right to remove access and implement additional fee (with advance notice) to view these data points. Contact your Lalamove Sales representative if you have any questions or feedback.")
	if err != nil {
		fmt.Println(err)
	}

	if err = f.SaveAs("Book4.xlsx"); err != nil {
		fmt.Println(err)
	}
}
