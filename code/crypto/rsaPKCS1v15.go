package main

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"log"
	_ "test/routers"
)

func main() {
	read:=rand.Reader
	privKey,err:=rsa.GenerateKey(read,1024)
	if err != nil{
		log.Fatalln(err.Error())
	}
	pubKey := privKey.PublicKey
	fmt.Println("N:",privKey.N)
	fmt.Println("公钥E:",privKey.E)
	fmt.Println("私钥D",privKey.D)

	plainText := "helloworld5645你好"

	ciphertext,err:=rsa.EncryptPKCS1v15(read,&pubKey,[]byte(plainText))
	if err != nil{
		log.Fatalln(err.Error())
	}
	fmt.Println(hexToString(ciphertext))

	result,err:=rsa.DecryptPKCS1v15(read,privKey,ciphertext)
	if err != nil{
		log.Fatalln(err.Error())
	}
	fmt.Println(string(result))

}

func hexToString(msg []byte)string{
	var result string
	for i := 0; i < len(msg); i++ {
		result += HEX2STRING[(msg[i]>>4)&0xf]
		result += HEX2STRING[msg[i]&0xf]
	}
	return result
}

var HEX2STRING  = []string{"0","1","2","3","4","5","6","7","8","9","a","b","c","d","e","f"}

