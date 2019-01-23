package main

import (
	"os"
	"strings"
	"path/filepath"
	"unicode"
	"strconv"
	"fmt"
	lua "github.com/yuin/gopher-lua"
	"github.com/360EntSecGroup-Skylar/excelize"
)

type Datarows struct {
	filename string
	sheetname string
	xlsx *excelize.File
}

const luaDatarowsTypeName = "datarows"

const numStyle =  `{"border":[{"type":"left","color":"000000","style":1},
		{"type":"top","color":"000000","style":1},
		{"type":"bottom","color":"000000","style":1},
		{"type":"right","color":"000000","style":1}],
		"number_format":4}`

const titleStyle = `{"border":[{"type":"left","color":"000000","style":1},
		{"type":"top","color":"000000","style":1},
		{"type":"bottom","color":"000000","style":1},
		{"type":"right","color":"000000","style":1}],
		"font":{"bold":true,"size":10},
		"alignment":{"horizontal":"center","vertical":"center"}}`
    

func registerDatarowsType(L *lua.LState) {
    mt := L.NewTypeMetatable(luaDatarowsTypeName)
    L.SetGlobal("datarows", mt)
    // static attributes
    L.SetField(mt, "new", L.NewFunction(newDatarows))
    // methods
    L.SetField(mt, "__index", L.SetFuncs(L.NewTable(), DatarowsMethods))
}

// Constructor
func newDatarows(L *lua.LState) int {
	filename := L.CheckString(1)
	sheetname := L.CheckString(2)
	xlsx := readExcel(filename,sheetname)
    dataRows := &Datarows{filename,sheetname,xlsx}
    ud := L.NewUserData()
    ud.Value = dataRows 
    L.SetMetatable(ud, L.GetTypeMetatable(luaDatarowsTypeName))
    L.Push(ud)
    return 1
}

// Checks whether the first lua argument is a *LUserData with *Person and returns this *Datarows.
func checkDatarows(L *lua.LState) *Datarows {
    ud := L.CheckUserData(1)
    if v, ok := ud.Value.(*Datarows); ok {
        return v
    }
    L.ArgError(1, "Datarows expected")
    return nil
}

var DatarowsMethods = map[string]lua.LGFunction{
    "get": datarowsGet,
    "set": datarowsSet,
    "save": datarowsSave,
}

// Getter for the Datarows row, col 
func datarowsGet(L *lua.LState) int {
	p := checkDatarows(L)
	axis := L.CheckString(2)
	value := p.xlsx.GetCellValue(p.sheetname, axis)
    L.Push(lua.LString(value))
    return 1
}

// Setter for the Datarows row, col 
func datarowsSet(L *lua.LState) int {
	p := checkDatarows(L)
	axis := L.CheckString(2)
	value := L.CheckString(3)
	idx := 0
	for i, r := range axis {
		if unicode.IsDigit(r) {
			idx = i
			break	
		}
	}

	// 没有合并单元格
	if idx == 1 {
		//数字样式
    	numstyle, err := p.xlsx.NewStyle(numStyle)
    	if err != nil {
		 	fmt.Println(err)
		} 
		if unicode.IsDigit(rune(value[0])) || strings.HasPrefix(value, "-") { 
			value1, _ := strconv.ParseFloat(value,64) 
			p.xlsx.SetCellValue(p.sheetname,axis,value1)
		}else {
			p.xlsx.SetCellValue(p.sheetname,axis,value)
		}
		p.xlsx.SetCellStyle(p.sheetname, axis, axis,numstyle)
	}else {
		//标题样式
    	titlestyle, err := p.xlsx.NewStyle(titleStyle)
    	if err != nil {
		 	fmt.Println(err)
		} 

		axisxStart := string(axis[0:1])+string(axis[idx:len(axis)])
		axisxEnd := string(axis[idx-1:len(axis)])
		p.xlsx.SetCellStyle(p.sheetname, axisxStart, axisxEnd, titlestyle)
		p.xlsx.MergeCell(p.sheetname, axisxStart , axisxEnd)
		p.xlsx.SetCellValue(p.sheetname, axisxStart,value)
	}
	return 1
}

// Setter for the Datarows row, col 
func datarowsSave(L *lua.LState) int {
	p := checkDatarows(L)
	if L.GetTop() == 2 {
		name := L.CheckString(2)
		p.xlsx.SaveAs(name)
        return 1
    }
	p.xlsx.Save()
    return 1
}

func readExcel(filename , sheetname string) (*excelize.File){
	xlsx, err := excelize.OpenFile(filename)
    if err != nil {
        fmt.Println(err)
        return nil
    }
	return xlsx
}


//独立起一个lua虚拟机
func testgo(path string) {
	L := lua.NewState()
	defer L.Close()

	registerDatarowsType(L)	

	if err := L.DoFile(`business.lua`); err != nil {
    	panic(err)
	}

	command := `computefile("`+path+`","平台展示")`
	if err := L.DoString(command); err != nil {
		fmt.Println("error " + path)
	}
}

func main() {
	L := lua.NewState()
	defer L.Close()

	registerDatarowsType(L)	

	if err := L.DoFile(`business.lua`); err != nil {
    	panic(err)
	}
	
	err := filepath.Walk("./files", func(path string, f os.FileInfo, err error) error {
		if ( f == nil ) {return err}
		if f.IsDir() {return nil}
		ok := strings.HasSuffix(path, ".xlsx")
  			if ok {
				command := `computefile("`+path+`","平台展示")`
				command = strings.Replace(command, "\\", "\\\\", -1)
				//fmt.Println(command)
				if err := L.DoString(command); err != nil {
					fmt.Println("error " + path)
					fmt.Println(err)
					return nil	
				}
				// 单独起一个lua虚拟机
				// go testgo(path)
			}
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
}