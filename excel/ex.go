package main

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func main() {
	xlsx, err := excelize.OpenFile("/Users/liweijian/code/py/project/exceltool/1_宏创价值成长1期-利润表_私募基金-20171220.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	// Get value from cell by given worksheet name and axis.
	cell := xlsx.GetCellValue("利润表_私募基金", "B2")
	fmt.Println(cell)
	// Get all the rows in the Sheet1.
	rows := xlsx.GetRows("利润表_私募基金")
	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}
}
