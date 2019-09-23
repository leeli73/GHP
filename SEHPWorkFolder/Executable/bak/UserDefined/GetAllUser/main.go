package main

import (
	"fmt"
	"strings"
    "database/sql"
    "log"
    _ "github.com/go-sql-driver/mysql"
)

// 获取表数据 
func Get(db *sql.DB) {
    rows, err := db.Query("select * from userinfo;")
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
		var UserName string
        for i, col := range values {
            if col == nil {
                value = "NULL"
            } else {
                value = string(col)
			}
			if strings.Compare(cloumns[i],"UserName") ==0{
				UserName = value
			}
        }
        if strings.Compare(UserName,"KEYNULL") != 0{
            fmt.Println(`<li class="list-group-item d-flex justify-content-between align-items-center">用户名:`+UserName+` 
					 	<span class="badge badge-primary badge-pill">
							 <a href="javascript:DeleteUser('`+UserName+`')" class="text-light">删除</a>
	  				 	</span>
	  				 </li>`)
        }
    }
    if err = rows.Err(); err != nil {
        log.Fatal(err)
    }
}
func main() {
	db, err := sql.Open("mysql", "root:root@/sehpassetshare")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	Get(db)
}