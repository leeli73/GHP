package main

import (
	"fmt"
	"flag"
	"strings"
    "database/sql"
    "log"
    _ "github.com/go-sql-driver/mysql"
)

func Insert(db *sql.DB,username string,passwd string,phone string,email string) {
    stmt, err := db.Prepare("INSERT INTO userinfo(UserName,UserPasswd,UserSubjection,UserPhone,UserEmail) VALUES(?,?,?,?,?);")
    if err != nil {
		fmt.Println("注册失败")
		return
    }
    _, err = stmt.Exec(username,passwd,"null",phone,email)
    if err != nil {
		fmt.Println("注册失败,用户已存在")
		return
    }
    fmt.Println("注册成功")
}
func main() {
	username := flag.String("username", "nil", "Input your username")
	passwd := flag.String("passwd", "nil", "Input your passwd")
	phone := flag.String("phone", "nil", "Input your phone")
	email := flag.String("email", "nil", "Input your email")
    flag.Parse()
	if strings.Compare(*username,"nil") == 0{
		fmt.Println("Error")
	} else {
		db, err := sql.Open("mysql", "root:root@/sehpassetshare")
		if err != nil {
			log.Fatal(err)
		}
        defer db.Close()
		Insert(db,*username,*passwd,*phone,*email)
	}
}