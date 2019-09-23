package main

import (
	"fmt"
	"flag"
	"strings"
    "database/sql"
    "net/url"
    "log"
    _ "github.com/go-sql-driver/mysql"
    "crypto/cipher"
    "crypto/aes"
    "bytes"
    "encoding/base64"
)

// 获取表数据 
func Get(db *sql.DB,id string) {
    rows, err := db.Query("select * from message where MessageTo='"+id+"';")
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
	i:=0
	outsign := -1
    for rows.Next() {
        err = rows.Scan(scanArgs...)
        if err != nil {
            log.Fatal(err)
		}
		var value string
		var MessageFrom string
		var MessageDate string
        for i, col := range values {
            if col == nil {
                value = "NULL"
            } else {
                value = string(col)
			}
			if strings.Compare(cloumns[i],"MessageFrom") ==0{
				MessageFrom = value
			} else if strings.Compare(cloumns[i],"MessageDate") ==0{
				MessageDate = value
			}
		}
			if outsign < 0{
				fmt.Println(`<a href="#" class="list-group-item list-group-item-action flex-column align-items-start active">
							<div class="d-flex w-100 justify-content-between">
							<h5 class="mb-1">`+MessageFrom+`</h5> <small>From</small>
							</div>
							<p class="mb-1">`+MessageDate+`</p> <small>To:`+id+`</small>
						</a>`)
			} else {
				fmt.Println(`<a href="#" class="list-group-item list-group-item-action flex-column align-items-start">
							<div class="d-flex w-100 justify-content-between">
							<h5 class="mb-1">`+MessageFrom+`</h5> <small class="text-muted">From</small>
							</div>
							<p class="mb-1">`+MessageDate+`</p> <small class="text-muted">To:`+id+`</small>
						</a>`)
			}
			outsign = outsign * -1
		i++
    }
    if i==0{
        fmt.Println(`<a href="#" class="list-group-item list-group-item-action flex-column align-items-start">
							<div class="d-flex w-100 justify-content-between">
							<h5 class="mb-1">尴尬了</h5> <small class="text-muted">From</small>
							</div>
							<p class="mb-1">你还没有留言呀</p> <small class="text-muted">To:`+id+`</small>
						</a>`)
    }
    if err = rows.Err(); err != nil {
        log.Fatal(err)
    }
}
func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
    padding := blockSize - len(ciphertext)%blockSize
    padtext := bytes.Repeat([]byte{byte(padding)}, padding)
    return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
    length := len(origData)
    unpadding := int(origData[length-1])
    return origData[:(length - unpadding)]
}

func AesEncrypt(origData, key []byte) ([]byte, error) {
    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }

    blockSize := block.BlockSize()
    origData = PKCS5Padding(origData, blockSize)
    blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
    crypted := make([]byte, len(origData))
    blockMode.CryptBlocks(crypted, origData)
    return crypted, nil
}

func AesDecrypt(crypted, key []byte) ([]byte, error) {
    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }

    blockSize := block.BlockSize()
    blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
    origData := make([]byte, len(crypted))
    blockMode.CryptBlocks(origData, crypted)
    origData = PKCS5UnPadding(origData)
    return origData, nil
}
func main() {
    username := flag.String("username", "nil", "Input your username")
	flag.Parse()
	if strings.Compare(*username,"nil") == 0{
		db, err := sql.Open("mysql", "root:root@/sehpassetshare")
		if err != nil {
			log.Fatal(err)
		}
        defer db.Close()
		Get(db,"Guest")
	} else {
		db, err := sql.Open("mysql", "root:root@/sehpassetshare")
		if err != nil {
			log.Fatal(err)
		}
        defer db.Close()
        aeskey := []byte("321423u9y8d2fwfl")
        decodeusername,err := url.QueryUnescape(*username)
        if err != nil {
            fmt.Println(err)
        }
        bytesPass, err := base64.StdEncoding.DecodeString(decodeusername)
        if err != nil {
            fmt.Println(err)
            return
        }

        tpass, err := AesDecrypt(bytesPass, aeskey)
        if err != nil {
            fmt.Println(err)
            return
        }
        usernameStr := string(tpass)
		Get(db,usernameStr)
	}
}