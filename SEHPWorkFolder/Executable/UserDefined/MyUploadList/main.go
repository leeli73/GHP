package main

import (
	"fmt"
    "flag"
    "net/url"
	"strings"
    "database/sql"
    "log"
    _ "github.com/go-sql-driver/mysql"
    "crypto/cipher"
    "crypto/aes"
    "bytes"
    "encoding/base64"
)

// 获取表数据 
func Get(db *sql.DB,id string) {
    rows, err := db.Query("select * from asset where AssetSubjection='"+id+"';")
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
    count := 0
    for rows.Next() {
        err = rows.Scan(scanArgs...)
        if err != nil {
            log.Fatal(err)
		}
		var value string
		var AssetID string
		var AssetName string
		var AssetURL string
        for i, col := range values {
            if col == nil {
                value = "NULL"
            } else {
                value = string(col)
			}
			if strings.Compare(cloumns[i],"AssetName") ==0{
				AssetName = value
			} else if strings.Compare(cloumns[i],"AssetURL") ==0{
				AssetURL = value
			} else if strings.Compare(cloumns[i],"AssetID") ==0{
				AssetID = value
			}
		}
		fmt.Println(`<li class="list-group-item d-flex justify-content-between align-items-center">`+AssetName+` 
					 	<span class="badge badge-primary badge-pill">
					 		<a href="`+AssetURL+`" class="text-light">下载</a>&nbsp;&nbsp; <a href="javascript:Delete('`+AssetID+`')" class="text-light">删除</a>
	  				 	</span>
                       </li>`)
        count ++
    }
    if err = rows.Err(); err != nil {
        log.Fatal(err)
    }
    if count == 0{
        fmt.Println(`<li class="list-group-item d-flex justify-content-between align-items-center">尴尬了...你还没有资源
					 	<span class="badge badge-primary badge-pill">
					 		<a href="javascript:alert('尴尬了...你还没有资源')" class="text-light">下载</a>&nbsp;&nbsp; <a href="javascript:alert('尴尬了...你还没有资源')" class="text-light">删除</a>
	  				 	</span>
	  				    </li>`)
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
		fmt.Println("Error")
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