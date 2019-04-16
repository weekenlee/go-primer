package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"unicode"

	"github.com/360EntSecGroup-Skylar/excelize"
	"liweijian.com/xls"
)

var jb_zjzh_dict map[string][]string
var zjzh_jb_dict map[string]string
var zjzh_khbh_dict map[string]string

//经办：{"file1":[[]], "file2":[[]]}
var item_dict map[string]map[string][][]string

func get_jb() {
	jb_zjzh_dict = make(map[string][]string)
	zjzh_jb_dict = make(map[string]string)
	if xlFile, err := xls.Open("资金账号与经办.xls", "utf-8"); err == nil {
		if sheet1 := xlFile.GetSheet(0); sheet1 != nil {
			for i := 0; i <= (int(sheet1.MaxRow)); i++ {
				row := sheet1.Row(i)
				for index := row.FirstCol(); index < row.LastCol(); index++ {
					if unicode.IsDigit(rune(row.Col(1)[0])) {
						v, ok := zjzh_khbh_dict[row.Col(1)]
						if ok {
							jb_zjzh_dict[row.Col(2)] = append(jb_zjzh_dict[row.Col(2)], v)
							zjzh_jb_dict[v] = row.Col(2)
						} else {
							//fmt.Println("找不到资金账户对应的客户编号,资金账号：" + row.Col(1))
						}
					}
				}
			}

		}
	}
}

func get_khbh() {
	zjzh_khbh_dict = make(map[string]string)
	if xlFile, err := xls.Open("广发对账单/20190415保证金.xls", "utf-8"); err == nil {
		if sheet1 := xlFile.GetSheet(0); sheet1 != nil {
			for i := 0; i <= (int(sheet1.MaxRow)); i++ {
				row := sheet1.Row(i)
				for index := row.FirstCol(); index < row.LastCol(); index++ {
					fmt.Println(row.Col(2), row.Col(0))
					zjzh_khbh_dict[row.Col(2)] = row.Col(0)
				}
			}

		}
	}
}

func handleFile(filepath string, fname string) {
	if strings.HasPrefix(fname, "资产托管部") {
		fmt.Println(filepath)

		if xlFile, err := xls.Open(filepath, "utf-8"); err == nil {
			if sheet1 := xlFile.GetSheet(0); sheet1 != nil {
				//寻找客户编号列
				row := sheet1.Row(0)
				keycol := -1
				var titlearr []string
				for index := row.FirstCol(); index < row.LastCol(); index++ {
					titlearr = append(titlearr, row.Col(index))
					if row.Col(index) == "客户编号" || row.Col(index) == "客户代码" {
						keycol = index
					}
				}

				//没有找到，不用处理这个文件
				if keycol == -1 {
					return
				}

				//找到客户编号,开始处理文件
				for i := 1; i <= (int(sheet1.MaxRow)); i++ {
					row := sheet1.Row(i)
					var arr []string
					for index := row.FirstCol(); index < row.LastCol(); index++ {
						//fmt.Printf("%v", row.Col(index))
						arr = append(arr, row.Col(index))
					}

					keyval := row.Col(keycol)
					jb := zjzh_jb_dict[keyval]
					v, ok := item_dict[jb]
					if ok {
						_, ok := v[fname]
						if ok {
							v[fname] = append(v[fname], arr)
						} else {
							var bigarr [][]string
							bigarr = append(bigarr, []string{filepath})
							bigarr = append(bigarr, titlearr)
							bigarr = append(bigarr, arr)
							v[fname] = bigarr
						}
					} else {
						item := make(map[string][][]string)
						var bigarr [][]string
						bigarr = append(bigarr, []string{filepath})
						bigarr = append(bigarr, titlearr)
						bigarr = append(bigarr, arr)
						item[fname] = bigarr
						item_dict[jb] = item
					}
				}
			}
		}
	}
}

func walkFilelist(path string) {
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
			//walkFilelist(path)
		}
		handleFile(path, f.Name())
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
}

func indexhelper(i, j int) string {
	chardict := "ABCDEFGHIJKLMNOPKRSTUVWXYZ"
	return string(chardict[j]) + strconv.Itoa(i+1)
}

func writeExcel(jb string, values map[string][][]string) {
	xlsx := excelize.NewFile()

	for k, v := range values {
		sheetname := strings.Split(k, "-")[1] + strings.Split(k, "-")[3]
		xlsx.NewSheet(sheetname)
		for i := 0; i < len(v); i++ {
			for j := 0; j < len(v[i]); j++ {
				idx := indexhelper(i, j)
				xlsx.SetCellValue(sheetname, idx, v[i][j])
			}
		}

	}
	xlsx.DeleteSheet("Sheet1")
	err := xlsx.SaveAs(jb + ".xlsx")
	if err != nil {
		fmt.Println(err)
	}

}

func main() {
	get_khbh()
	return
	get_jb()
	//fmt.Println(jb_zjzh_dict)
	//fmt.Println(zjzh_jb_dict)

	item_dict = make(map[string]map[string][][]string)
	walkFilelist("广发对账单")

	for k, v := range item_dict {
		// fmt.Println(k)
		// for kk, vv := range v {
		// fmt.Println(kk)
		// fmt.Println(vv)
		// }
		writeExcel(k, v)
	}
}
