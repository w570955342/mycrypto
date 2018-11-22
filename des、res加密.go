package main

import (
	"fmt"
	"mycrypto/aes"
	"mycrypto/des"
)

func main() {
lable:
	fmt.Println("请选择模式：DES加密按1  DES解密按2  AES加密按3  AES解密按4 退出ctr+c！")
	var mod string
	var arr []rune
	var a rune
	for i := 0; ; i++ {
		fmt.Scanf("%c", &a)
		arr = append(arr, a)
		//arr[i] = a err  必须用append增加元素，arr本身长度、容量均为0
		if arr[i] == 10 {
			break
		}
	}
	//兼容Windows的.exe和goland  Windows下的.exe 回车键代表\r\n 两个字节 [13 10] goland下回车键代表\n [10]
	//必须根据倒数第二个字符\r来判断 不能先把切片转化成字符串 然后mod[:1] 只截取第一位 这样导致12和1是一样的效果
	if arr[len(arr)-2] == 13 {
		mod = string(arr[:len(arr)-2])
	} else {
		mod = string(arr[:len(arr)-1])
	}
	if mod == "1" {
		des.CbcEncryptFileAll()
	} else if mod == "2" {
		des.CbcDecryptFileAll()
	} else if mod == "3" {
		aes.CtrEncryptFileAll()
	} else if mod == "4" {
		aes.CtrDecryptFileAll()
	} else if mod == "5" {
		fmt.Println("暂未开启该功能，敬请期待！")
	} else if mod == "6" {
		fmt.Println("暂未开启该功能，敬请期待！")
	} else {
		fmt.Println("你是不是傻！")
	}
	goto lable
}
