package main

import (
	"fmt"
	"mycrypto/rsa"
)
func main() {
	rsa.GenerateKey(512)
	src:=[]byte("好好学习天天向上")
	cipherText:= rsa.Encrypt(src,"public.pem")
	fmt.Println(string(cipherText))
	plainText:= rsa.Decrypt(cipherText,"private.pem")
	fmt.Println(string(plainText))
}
