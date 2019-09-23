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
    rows, err := db.Query("select * from message;")
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
		var MessageID string
		var MessageFrom string
		var MessageTo string
		var Message string
        for i, col := range values {
            if col == nil {
                value = "NULL"
            } else {
                value = string(col)
			}
			if strings.Compare(cloumns[i],"MessageID") ==0{
				MessageID = value
			} else if strings.Compare(cloumns[i],"MessageFrom") ==0{
				MessageFrom = value
			} else if strings.Compare(cloumns[i],"MessageTo") ==0{
				MessageTo = value
			} else if strings.Compare(cloumns[i],"MessageDate") ==0{
				Message = value
			}
		}
		fmt.Println(`<li class="list-group-item d-flex justify-content-between align-items-center">From:`+MessageFrom+`,To:`+MessageTo+`,Data:`+Message+` 
					 	<span class="badge badge-primary badge-pill">
							 <a href="javascript:DeleteMessage('`+MessageID+`')" class="text-light">删除</a>
	  				 	</span>
	  				 </li>`)
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