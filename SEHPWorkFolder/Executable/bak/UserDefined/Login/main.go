package main

import (
	"fmt"
	"flag"
	"strings"
    "database/sql"
    "log"
    _ "github.com/go-sql-driver/mysql"
)
func Get(db *sql.DB,username string) string{
    rows, err := db.Query("select UserPasswd from userinfo where UserName='"+username+"';")
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()
    cloumns, err := rows.Columns()  
    if err != nil {
        log.Fatal(err)
    }
    values := make([]sql.RawBytes, len(cloumns))
    scanArgs := make([]interface{}, len(values))
    for i := range values {
        scanArgs[i] = &values[i]
    }
    for rows.Next() {
        err = rows.Scan(scanArgs...)
        if err != nil {
            log.Fatal(err)
        }
        var value string
        for _, col := range values {
            if col == nil {
                value = "NULL"
            } else {
                value = string(col)
            }
			return value
        }
    }
    if err = rows.Err(); err != nil {
        log.Fatal(err)
	}
	return "Error"
}
func main() {
	username := flag.String("username", "nil", "Input your username")
	passwd := flag.String("passwd", "nil", "Input your passwd")
    flag.Parse()
	if strings.Compare(*username,"nil") == 0 || strings.Compare(*passwd,"nil") == 0{
		fmt.Println("Error")
	} else {
		db, err := sql.Open("mysql", "root:root@/sehpassetshare")
		if err != nil {
			fmt.Println("Error")
		}
        defer db.Close()
		pass := Get(db,*username)
		if strings.Compare(pass,*passwd) == 0{
			fmt.Print("登录成功")
		} else {
			fmt.Print("登录失败")
		}
	}
}