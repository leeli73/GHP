package main

import (
    "database/sql"
	"log"
	"flag"
    "crypto/md5"
    "strconv"
    "time"
    "math/rand"
    _ "github.com/go-sql-driver/mysql"
    "encoding/hex"
)
func Insert(mydb *sql.DB,id string,name string,username string,public string,url string) {
    stmt, err := mydb.Prepare("INSERT INTO asset(AssetID,AssetURL,AssetSubjection,isPublic,AssetName) VALUES(?,?,?,?,?);")
    if err != nil {
		return
    }
    _, err = stmt.Exec(id,url,username,public,name)
    if err != nil {
		return
    }
}
func main(){
	name := flag.String("name", "nil", "Input your username")
	username := flag.String("username", "nil", "Input your username")
	public := flag.String("public", "0", "Input your username")
    flag.Parse()
	hash := md5.New()
	hashStr := strconv.FormatInt(time.Now().Unix()+int64(rand.Intn(100))*int64(rand.Intn(10)),10)
	hash.Write([]byte(hashStr))
	cipherStr := hash.Sum(nil)
    ID := hex.EncodeToString(cipherStr)
	URL := "http://127.0.0.1:88/asset/"+*username+"_"+*name
	db, err := sql.Open("mysql", "root:root@/sehpassetshare")
	if err != nil {
		log.Fatal(err)
	}
    Insert(db,ID,*name,*username,*public,URL)
}