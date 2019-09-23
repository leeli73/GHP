package main

import (
    "fmt"
    "strings"
    "net/url"
    "flag"
    "database/sql"
    "log"
    _ "github.com/go-sql-driver/mysql"
    "crypto/cipher"
    "crypto/aes"
    "bytes"
    "encoding/base64"
)
// 更新数据
func UpdatePasswd(db *sql.DB,username string,passwd string) {
    stmt, err := db.Prepare("UPDATE userinfo SET UserPasswd='"+passwd+"' WHERE UserName='"+username+"';")
    if err != nil {
        fmt.Println("更新失败")
        return 
    }
    _, err = stmt.Exec()
    if err != nil{
        fmt.Println("更新失败")
        return 
    }
    fmt.Println("更新成功")
    
}
func UpdateSub(db *sql.DB,username string,sub string) {
    stmt, err := db.Prepare("UPDATE userinfo SET UserSubjection='"+sub+"' WHERE UserName='"+username+"';")
    if err != nil {
        fmt.Println("更新失败")
        return 
    }
    _, err = stmt.Exec()
    if err != nil{
        fmt.Println("更新失败")
        return 
    }
    fmt.Println("更新成功")
    
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
    passwd := flag.String("passwd", "nil", "Input your passwd")
    sub := flag.String("sub", "nil", "Input your sub")
    flag.Parse()
    if strings.Compare(*username,"nil") == 0{
		fmt.Println("Error")
	} else {
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
        db, err := sql.Open("mysql", "root:root@/sehpassetshare")
        if err != nil {
            log.Fatal(err)
        }
        defer db.Close()
        if strings.Compare(*passwd,"nil") == 0{
            UpdateSub(db,usernameStr,*sub)
        } else if strings.Compare(*sub,"nil") == 0{
            UpdatePasswd(db,usernameStr,*passwd)
        }
    }
}