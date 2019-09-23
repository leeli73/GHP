package main

import (
	"fmt"
	"flag"
	"strings"
	"net/url"
    "database/sql"
    "log"
	_ "github.com/go-sql-driver/mysql"
	"crypto/md5"
    "strconv"
    "time"
	"math/rand"
	"encoding/hex"
)

func Insert(db *sql.DB,MessageID string,MessageFrom string,MessageTo string,MessageDate string) {
    stmt, err := db.Prepare("INSERT INTO message(MessageID,MessageFrom,MessageTo,MessageDate) VALUES(?,?,?,?);")
    if err != nil {
		fmt.Println("留言失败")
		return
    }
    _, err = stmt.Exec(MessageID,MessageFrom,MessageTo,MessageDate)
    if err != nil {
		fmt.Println("留言失败")
		return
    }
    fmt.Println("留言成功")
}
func main() {
	MessageFrom := flag.String("MessageFrom", "nil", "Input your MessageFrom")
	MessageTo := flag.String("MessageTo", "nil", "Input your MessageTo")
	MessageDate := flag.String("MessageDate", "nil", "Input your MessageDate")
    flag.Parse()
	if strings.Compare(*MessageFrom,"nil") == 0 || strings.Compare(*MessageTo,"nil") == 0 || strings.Compare(*MessageDate,"nil") == 0{
		fmt.Println("Error")
	} else {
		db, err := sql.Open("mysql", "root:root@/sehpassetshare")
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()
		MessageFromStr,err := url.QueryUnescape(*MessageFrom)
        if err != nil {
            fmt.Println(err)
		}
		MessageToStr,err := url.QueryUnescape(*MessageTo)
        if err != nil {
            fmt.Println(err)
		}
		MessageDateStr,err := url.QueryUnescape(*MessageDate)
        if err != nil {
            fmt.Println(err)
		}
		hash := md5.New()
		hashStr := strconv.FormatInt(time.Now().Unix()+int64(rand.Intn(100))*int64(rand.Intn(10)),10)
		hash.Write([]byte(hashStr))
		cipherStr := hash.Sum(nil)
		ID := hex.EncodeToString(cipherStr)
		Insert(db,ID,MessageFromStr,MessageToStr,MessageDateStr)
	}
}