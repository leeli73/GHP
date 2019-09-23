package main

import (
    "fmt"
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
        for i, col := range values {
            if col == nil {
                value = "NULL"
            } else {
                value = string(col)
            }
            fmt.Println(cloumns[i], ": ", value)
        }
        fmt.Println("------------------")
    }
    if err = rows.Err(); err != nil {
        log.Fatal(err)
    }
}

// 插入数据
func Insert(db *sql.DB) {
    stmt, err := db.Prepare("INSERT INTO user(name, age) VALUES(?, ?);")
    if err != nil {
        log.Fatal(err)
    }
    res, err := stmt.Exec("python", 19)
    if err != nil {
        log.Fatal(err)
    }
    lastId, err := res.LastInsertId()
    if err != nil {
        log.Fatal(err)
    }
    rowCnt, err := res.RowsAffected()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("ID=%d, affected=%d\n", lastId, rowCnt)
}

// 删除数据
func Delete(db *sql.DB) {
    stmt, err := db.Prepare("DELETE FROM user WHERE name='python'")
    if err != nil {
        log.Fatal(err)
    }
    res, err := stmt.Exec()
    lastId, err := res.LastInsertId()
    if err != nil {
        log.Fatal(err)
    }
    rowCnt, err := res.RowsAffected()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("ID=%d, affected=%d\n", lastId, rowCnt)
}

// 更新数据
func Update(db *sql.DB) {
    stmt, err := db.Prepare("UPDATE user SET age=27 WHERE name='python'")
    if err != nil {
        log.Fatal(err)
    }
    res, err := stmt.Exec()
    lastId, err := res.LastInsertId()
    if err != nil {
        log.Fatal(err)
    }
    rowCnt, err := res.RowsAffected()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("ID=%d, affected=%d\n", lastId, rowCnt)
}

func main() {
    db, err := sql.Open("mysql", "root:root@/sehpassetshare")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
    Get(db)
}