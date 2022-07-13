package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"

)

func main() {
	f := excelize.NewFile()
	// Create a new sheet.
	index := f.NewSheet("Sheet1")
	// Set value of a cell.
	//f.SetCellValue("Sheet2", "A2", "Hello world.")
	f.SetCellValue("Sheet1", "B2", 100)
	// Set active sheet of the workbook.
	f.SetActiveSheet(index)
	// Save spreadsheet by the given path.

	//basicStyle, err := f.NewStyle(&excelize.Style{
	//	Fill: excelize.Fill{Type: "center", Pattern: 0, Color: []string{"FAFAD2"}, Shading: 0},
	//	Border: []excelize.Border{
	//		{Type: "left", Color: "0000FF", Style: 3},
	//	},
	//})

	basicStyle, err := f.NewStyle(&excelize.Style{
		Fill: excelize.Fill{Type: "gradient", Color: []string{"#FFFFFF", "#E0EBF5"}, Shading: 1},
	})

	basicStyle, err = f.NewStyle(`{"fill":{"type":"pattern","color":["#FFFAF2"],"pattern":1}}`)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetCellStyle("Sheet1", "A2", "A2", basicStyle)

	styleID1, err := f.NewStyle(`{"fill":{"type":"pattern","pattern":1,"color":["#DC143C"]}}`)

	err = f.SetCellStyle("Sheet1", "B4", "B4", styleID1)

	if err = f.SaveAs("Book3.xlsx"); err != nil {
		fmt.Println(err)
	}
}
