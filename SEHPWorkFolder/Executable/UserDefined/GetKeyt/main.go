package main

import (
	"fmt"
	"strings"
	"flag"
    "crypto/cipher"
    "crypto/aes"
    "bytes"
    "encoding/base64"
)

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
	username := flag.String("username", "mil", "Input your username")
	flag.Parse()
	if strings.Compare(*username,"nil") != 0{
		var aeskey = []byte("321423u9y8d2fwfl")
		pass := []byte(*username)
		xpass, err := AesEncrypt(pass, aeskey)
		if err != nil {
			fmt.Println(err)
			return
		}
		pass64 := base64.StdEncoding.EncodeToString(xpass)
    	fmt.Printf("%v\n",pass64)
	}
    /*var aeskey = []byte("321423u9y8d2fwfl")
    pass := []byte("vdncloud123456")
    xpass, err := AesEncrypt(pass, aeskey)
    if err != nil {
        fmt.Println(err)
        return
    }

    pass64 := base64.StdEncoding.EncodeToString(xpass)
    fmt.Printf("加密后:%v\n",pass64)

    bytesPass, err := base64.StdEncoding.DecodeString(pass64)
    if err != nil {
        fmt.Println(err)
        return
    }

    tpass, err := AesDecrypt(bytesPass, aeskey)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Printf("解密后:%s\n", tpass)*/
}