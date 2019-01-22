package main

import (
           "database/sql"
           _"github.com/mattn/go-oci8"
           "log"
           "os"
)

func main() {
           log.SetFlags(log.Lshortfile| log.LstdFlags)
           log.Println("OracleDriver example")
           os.Setenv("NLS_LANG","")
           // 用户名/密码@实例名 跟sqlplus的conn命令类似
           db,err := sql.Open("oci8", "hsfa/hsfa123@10.35.68.112:1521/zctgdb")
           //db,err := sql.Open("oci8", "hsfa/hsfa_0601@172.19.1.208:1521/WBGZDB")
           if err != nil {
                     log.Fatal(err)
           }
           rows,err := db.Query("select 3.14, 'foo' from dual")
           if err != nil {
                     log.Fatal(err)
           }
           defer db.Close()
           for rows.Next() {
                     var f1 float64
                     var f2 string
                     rows.Scan(&f1,&f2)
                     log.Println(f1,f2) // 3.14 foo
           }
           rows.Close()
 
}
