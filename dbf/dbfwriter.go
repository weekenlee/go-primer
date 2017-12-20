package main

import (
	"bufio"
	"container/list"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"liweijian/xls"

	"github.com/LindsayBradford/go-dbf/godbf"
	lua "github.com/yuin/gopher-lua"
)

//银行存款表
var gmap map[string]string

//子基金代码表
var gsonsmmap map[string]string

//日志类
var debugLog *log.Logger

//由lua配置的路径
var rootdir string
var bankexcelname string
var sonsmexcelname string
var todir string

//记录
type record struct {
	Custname string
	D_jyrq   string
	L_ywlb   string
	Vc_zjzh  string
	L_sclb   string
	Vc_isin  string
	L_zqlb   string
	En_cjje  string
	En_cjjg  string
	L_cjsl   string
	En_fy    string
}

var listrecord *list.List

//初始化dbf文件
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

//dbf文件增加一条记录
func addRow(row1 *xls.Row, ywdm string, bankacc string, day string, sonsm string) {
	je, _ := strconv.ParseFloat(row1.Col(9), 64)
	sl, _ := strconv.ParseFloat(row1.Col(10), 64)
	dj := je / sl
	djstr := fmt.Sprintf("%.2f", dj)
	r := record{
		Custname: row1.Col(1),
		D_jyrq:   day,
		L_ywlb:   ywdm,
		Vc_zjzh:  bankacc,
		L_sclb:   "4",
		Vc_isin:  sonsm,
		L_zqlb:   "11",
		En_cjje:  row1.Col(10),
		En_cjjg:  djstr,
		L_cjsl:   row1.Col(11),
		En_fy:    row1.Col(12)}
	listrecord.PushBack(r)
}

//初始化银行存款map
func initBankExcel(filename string) {
	if xlFile, err := xls.Open(filename, "utf-8"); err == nil {
		if sheet1 := xlFile.GetSheet(0); sheet1 != nil {
			for i := 4; i <= (int(sheet1.MaxRow)); i++ {
				row1 := sheet1.Row(i)
				gmap[row1.Col(3)] = row1.Col(4)
			}

		}
	}
}

//查找银行存款账号
func getBankAccount(tacode string) (bankacc string, e error) {
	if gmap == nil {
		gmap = make(map[string]string, 500)
		initBankExcel(bankexcelname)
	}
	if v, ok := gmap[tacode]; ok {
		return v, nil
	}
	return "", errors.New("not Found:" + tacode)
}

//初始化子基金代码map
func initSonsmExcel(filename string) {
	if xlFile, err := xls.Open(filename, "utf-8"); err == nil {
		if sheet1 := xlFile.GetSheet(0); sheet1 != nil {
			for i := 1; i <= (int(sheet1.MaxRow)); i++ {
				row1 := sheet1.Row(i)
				gsonsmmap[strings.TrimSpace(row1.Col(6))] = row1.Col(5)
			}

		}
	}
}

//查找子基金代码
func getSonsm(name string) (sonsm string, e error) {
	if gsonsmmap == nil {
		gsonsmmap = make(map[string]string, 500)
		initSonsmExcel(sonsmexcelname)
	}
	if v, ok := gsonsmmap[strings.TrimSpace(name)]; ok {
		return v, nil
	}
	return "", errors.New("not Found:" + name)
}

func splitList() {
	mapArray := make(map[string]*list.List)
	for v := listrecord.Front(); v != nil; v = v.Next() {
		// fmt.Println(mapArray)
		re := v.Value.(record)
		li := mapArray[re.Custname+re.D_jyrq]
		if li == nil {
			li = list.New()
			mapArray[re.Custname+re.D_jyrq] = li
		}
		li.PushBack(re)
	}

	for k, v := range mapArray {
		// fmt.Println("处理" + k)
		debugLog.Println("处理文件:" + k)
		writeRecord(v)
	}
}

func writeRecord(records *list.List) {
	dbfTable := initdbf()
	for v := records.Front(); v != nil; v = v.Next() {
		rownum := dbfTable.NumberOfRecords()
		dbfTable.AddNewRecord()
		re := v.Value.(record)
		dbfTable.SetFieldValueByName(rownum, "D_JYRQ", re.D_jyrq)
		dbfTable.SetFieldValueByName(rownum, "l_BH", strconv.Itoa(rownum))
		dbfTable.SetFieldValueByName(rownum, "l_YWLB", re.L_ywlb)
		dbfTable.SetFieldValueByName(rownum, "VC_ZJZH", re.Vc_zjzh)
		dbfTable.SetFieldValueByName(rownum, "L_SCLB", re.L_sclb)
		dbfTable.SetFieldValueByName(rownum, "VC_ISIN", re.Vc_isin)
		dbfTable.SetFieldValueByName(rownum, "L_ZQLB", re.L_zqlb)
		dbfTable.SetFieldValueByName(rownum, "EN_CJJE", re.En_cjje)
		dbfTable.SetFieldValueByName(rownum, "EN_CJJG", re.En_cjjg)
		dbfTable.SetFieldValueByName(rownum, "L_CJSL", re.L_cjsl)
		dbfTable.SetFieldValueByName(rownum, "EN_FY", re.En_fy)
		dbfTable.SetFieldValueByName(rownum, "L_GLBH", strconv.Itoa(rownum))

		debugLog.Printf("%s", "写入记录: ")
		t := reflect.TypeOf(re)
		v := reflect.ValueOf(re)
		for k := 0; k < t.NumField(); k++ {
			debugLog.Printf(" %s-%v", t.Field(k).Name, v.Field(k).Interface())
			fmt.Printf(" %s-%v", t.Field(k).Name, v.Field(k).Interface())
		}
		debugLog.Println("")
		fmt.Println("")
	}
	re := records.Front().Value.(record)
	fmt.Println("写入 " + re.Custname)
	filename := "TGTTRADEDATA_" + re.D_jyrq + ".dbf"
	dir := todir //os.Getwd()
	dirpath := dir + "/data/" + re.Custname + "/" + re.D_jyrq
	filepath := dirpath + "/" + filename

	err := os.MkdirAll(dirpath, os.ModePerm) //生成多级目录
	if err != nil {
		fmt.Println(err)
	}
	os.Remove(filepath)
	dbfTable.SaveFile(filepath)
}

func handleFile(filename string) {
	if xlFile, err := xls.Open(filename, "utf-8"); err == nil {
		if sheet1 := xlFile.GetSheet(0); sheet1 != nil {
			for i := 4; i <= (int(sheet1.MaxRow)); i++ {
				row1 := sheet1.Row(i)
				col1 := row1.Col(0)
				if len(strings.TrimSpace(col1)) != 0 {
					if strings.IndexAny(strings.TrimSpace(col1), "总计") == 0 {
						break
					}
					//非机构客户，忽略
					if len(row1.Col(2)) == 0 {
						continue
					}
					//查找银行存款账户
					bankacc, err := getBankAccount(row1.Col(2))
					if err != nil {
						fmt.Println("找不到银行账户，产品代码：", row1.Col(2))
						continue
					}
					sonsmcode := ""
					sonsmcode, err = getSonsm(row1.Col(7))
					if err != nil {
						fmt.Println("！！！！！！！！！无法找到 " + row1.Col(7))
						sonsmcode = "待补充" + row1.Col(7)
					}
					if strings.Contains(row1.Col(4), "申购") {
						addRow(row1, "1901", bankacc, row1.Col(5), sonsmcode) // 申请用申请日期
						addRow(row1, "1902", bankacc, row1.Col(6), sonsmcode) // 确认用确认日期
					} else if strings.Contains(row1.Col(4), "认购") {
						addRow(row1, "1913", bankacc, row1.Col(5), sonsmcode) // 申请用申请日期
						addRow(row1, "1915", bankacc, row1.Col(6), sonsmcode) // 确认用确认日期
					} else if strings.Contains(row1.Col(4), "赎回") {
						addRow(row1, "1903", bankacc, row1.Col(6), sonsmcode)
						addRow(row1, "1904", bankacc, row1.Col(6), sonsmcode)
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
		match, _ := regexp.MatchString(".*基金交易确认明细表.*.xls$", f.Name())
		if match {
			handleFile(path)
		}
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
}

func end() {
	fmt.Println("执行完毕")
	reader := bufio.NewReader(os.Stdin)
	msg, _ := reader.ReadString('\n')
	fmt.Printf(msg)
}

func initconfig() {
	L := lua.NewState()
	defer L.Close()
	if err := L.DoFile("config.lua"); err != nil {
		panic(err)
	}
	if err := L.CallByParam(lua.P{
		Fn:      L.GetGlobal("getconfig"),
		NRet:    4,
		Protect: true,
	}); err != nil {
		panic(err)
	}
	rootdir = string(L.Get(-4).(lua.LString))        // returned value
	sonsmexcelname = string(L.Get(-3).(lua.LString)) // returned value
	bankexcelname = string(L.Get(-2).(lua.LString))  // returned value
	todir = string(L.Get(-1).(lua.LString))
	L.Pop(4)
}

func main() {
	initconfig()
	fileName := "Info_" + time.Now().Format("2006-01-02150405") + ".log"
	logFile, err := os.Create(fileName)
	if err != nil {
		log.Fatalln("open file error")
	}
	debugLog = log.New(logFile, "[Info]", log.LstdFlags)

	listrecord = list.New()

	walkFilelist(rootdir)
	splitList()
	end()
}
