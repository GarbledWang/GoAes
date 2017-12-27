package main

import (
	"flag"
	"fmt"
	"os"
	"bufio"
	"io"
	"encoding/base64"
)

func main() {
	var password string
	var filePath string
	var aesType int
	flag.StringVar(&password, "password", "", "Your Password")
	flag.StringVar(&filePath, "filePath", "", "File Absolute Path")
	flag.IntVar(&aesType, "type", 0, "Enrypt 1,Decrypt 2")
	flag.Parse() //解析输入的参数
	if password == "" {
		fmt.Println("密码为null")
		return
	}
	if filePath == "" {
		fmt.Println("地址为null")
		return
	}
	//创建Aes
	aes := AesEncrypt{}
	key := aes.GetKey(password)
	//读取文件
	content := readFile(filePath)
	if aesType == 1 {
		v, err := aes.Encrypt(content, key)
		check(err)
		fileCreate([]byte(base64.StdEncoding.EncodeToString(v)), "entrypt.txt")
		fmt.Println(string(v))
	} else if aesType == 2 {
		//解密
		bytes,err := base64.StdEncoding.DecodeString(content)
		check(err)
		v, err := aes.Decrypt(bytes, key)
		check(err)
		fileCreate([]byte(v), "decrypt.txt")
		fmt.Println(v)
	}

}
func fileCreate(v []byte, fileName string) {
	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModeAppend|os.ModePerm)
	defer f.Close()
	check(err)
	f.WriteString(string(v))
}

func readFile(filePath string) string {
	file, err := os.Open(filePath)
	defer file.Close()
	check(err)
	reader := bufio.NewReader(file)
	var content string = ""
	for {
		a, _, c := reader.ReadLine()
		if c == io.EOF {
			break
		}
		content += string(a)
	}
	return content
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
