package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"liweijian.com/xls"

	"github.com/LindsayBradford/go-dbf/godbf"
)

var datestr string

func initdbf() *godbf.DbfTable {
	dbfTable := godbf.New("gbk")
	dbfTable.AddDateField("D_JYRQ")
	dbfTable.AddNumberField("l_BH", 10, 0)
	dbfTable.AddNumberField("l_YWLB", 4, 0)
	dbfTable.AddTextField("VC_ZJZH", 32)
	dbfTable.AddNumberField("L_SCLB", 1, 0)
	dbfTable.AddTextField("VC_ISIN", 20)
	dbfTable.AddNumberField("L_ZQLB", 4, 0)
	dbfTable.AddNumberField("EN_CJJE", 16, 2)
	dbfTable.AddNumberField("EN_CJJG", 14, 8)
	dbfTable.AddNumberField("L_CJSL", 14, 2)
	dbfTable.AddNumberField("EN_FY", 12, 2)
	dbfTable.AddNumberField("L_GLBH", 10, 0)
	return dbfTable
}

func addRow(row1 *xls.Row, dbfTable *godbf.DbfTable, ywdm string) {
	rownum := dbfTable.NumberOfRecords()
	dbfTable.AddNewRecord()
	dbfTable.SetFieldValueByName(rownum, "D_JYRQ", row1.Col(5))
	dbfTable.SetFieldValueByName(rownum, "l_BH", strconv.Itoa(rownum))
	dbfTable.SetFieldValueByName(rownum, "l_YWLB", ywdm)
	dbfTable.SetFieldValueByName(rownum, "VC_ZJZH", "虚拟")
	dbfTable.SetFieldValueByName(rownum, "L_SCLB", "4")
	dbfTable.SetFieldValueByName(rownum, "VC_ISIN", "SM0792")
	dbfTable.SetFieldValueByName(rownum, "L_ZQLB", "11")
	dbfTable.SetFieldValueByName(rownum, "EN_CJJE", row1.Col(9))
	dbfTable.SetFieldValueByName(rownum, "EN_CJJG", "11.2")
	dbfTable.SetFieldValueByName(rownum, "L_CJSL", row1.Col(10))
	dbfTable.SetFieldValueByName(rownum, "EN_FY", row1.Col(11))
	dbfTable.SetFieldValueByName(rownum, "L_GLBH", strconv.Itoa(rownum))
	datestr = row1.Col(5)
}

func handleFile(filename string, dbfTable *godbf.DbfTable) {
	fmt.Println("handle " + filename)
	if xlFile, err := xls.Open(filename, "utf-8"); err == nil {
		if sheet1 := xlFile.GetSheet(0); sheet1 != nil {
			fmt.Println("Total Lines ", sheet1.MaxRow)
			for i := 4; i <= (int(sheet1.MaxRow)); i++ {
				row1 := sheet1.Row(i)
				col1 := row1.Col(0)
				if len(strings.TrimSpace(col1)) != 0 {
					if strings.IndexAny(strings.TrimSpace(col1), "总计") == 0 {
						break
					}
					fmt.Println(row1.Col(3))
					if strings.Contains(row1.Col(3), "申购") {
						addRow(row1, dbfTable, "1901")
						addRow(row1, dbfTable, "1902")
					} else if strings.Contains(row1.Col(3), "认购") {
						addRow(row1, dbfTable, "1913")
						addRow(row1, dbfTable, "1915")
					} else if strings.Contains(row1.Col(3), "赎回") {
						addRow(row1, dbfTable, "1903")
						addRow(row1, dbfTable, "1904")
					}
				}
			}
		}
	}
}

func walkFilelist(path string, dbfTable *godbf.DbfTable) {
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		fmt.Println(f.Name())
		match, _ := regexp.MatchString(".*基金交易确认明细表.*.xls$", f.Name())
		if match {
			handleFile(f.Name(), dbfTable)
		}
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
}

func main() {
	dbfTable := initdbf()
	walkFilelist("./", dbfTable)
	os.Remove("TGTTRADEDATA_" + datestr + ".dbf")
	dbfTable.SaveFile("TGTTRADEDATA_" + datestr + ".dbf")
}
