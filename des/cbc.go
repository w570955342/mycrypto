package des

import (
	"fmt"
	"io"
	"mycrypto/cipher"
	"os"
)

//封装加密拷贝文件函数
func CbcEncryptFile(src, des string) {
	f1, err1 := os.Open(src)

	if err1 != nil {
		fmt.Println("os.Open err:", err1)
		return
	}

	defer f1.Close()

	f2, err2 := os.Create(des)
	if err2 != nil {
		fmt.Println("os.Create err:", err2)
		return
	}

	defer f2.Close()

	//buf := make([]byte, 4096)
	key := []byte("1234abdd")
	for {
		buf := make([]byte, 4096*4096)
		n, err3 := f1.Read(buf)
		if err3 != nil && err3 != io.EOF {
			fmt.Println("f.Read err:", err3)
			return
		}
		if n == 0 {
			//fmt.Println("被加密文件读取完毕！")
			return
		}

		//fmt.Printf("加密前数据长度：%d\n",len(buf))
		buf = cipher.DesCbcEncrypt(buf[:n], key)
		_, err4 := f2.Write(buf)
		if err4 != nil {
			fmt.Println("f2.Write err:", err4)
			return
		}
	}
}

//封装解密拷贝文件函数
func CbcDecryptFile(src, des string) {

	f1, err1 := os.Open(src)
	if err1 != nil {
		fmt.Println("os.Open err:", err1)
		return
	}

	defer f1.Close()

	f2, err2 := os.Create(des)
	if err2 != nil {
		fmt.Println("os.Create err:", err2)
		return
	}

	defer f2.Close()

	//buf := make([]byte, 4096)
	key := []byte("1234abdd")
	for {
		buf := make([]byte, 4096*4096+8)
		n, err3 := f1.Read(buf)
		if err3 != nil && err3 != io.EOF {
			fmt.Println("f.Read err:", err3)
			return
		}
		if n == 0 {
			//fmt.Println("被解密文件读取完毕！")
			return
		}

		buf = cipher.DesCbcDecrypt(buf[:n], key)
		_, err4 := f2.Write(buf)
		if err4 != nil {
			fmt.Println("f2.Write err:", err4)
			return
		}
	}
}


//对路径的读取进一步封装
func CbcEncryptFileAll() {
	os.MkdirAll("./加密/", 0777)
	fmt.Println("请输入要DES加密的文件的路径：")
lable:
	//获取路径，兼容路径中含有空格
	var src string
	var arr []rune
	var a rune

	for i := 0; ; i++ {
		fmt.Scanf("%c", &a)
		arr = append(arr, a)
		//arr[i] = a err  必须用append增加元素，arr本身长度、容量均为0
		if arr[i] == 10 {
			fmt.Print(string(arr))
			break
		}
	}
	//兼容Windows的.exe和goland  Windows下的.exe 回车键代表\r\n 两个字节 [13 10] goland下回车键代表\n [10]
	if arr[len(arr)-2] == 13 {
		if arr[0]==34 {
			src = string(arr[1:len(arr)-3])
		}else {
			src = string(arr[:len(arr)-2])
		}
	} else {
		//为了在goland中运行 goland中回车键代表\n
		src = string(arr[:len(arr)-1])
	}

	f, err := os.Open(src)
	if err != nil {
		fmt.Println("输入路径有误！请重新输入！")
		goto lable
	}
	defer f.Close()
	fileInfo, err := f.Stat()
	fmt.Println("文件正在进行DES加密，请稍安勿躁……………")
	CbcEncryptFile(src, "./加密/DE"+fileInfo.Name())
	fmt.Println("文件DES加密成功！")
}

func CbcDecryptFileAll() {

	os.MkdirAll("./解密/", 0777)
	fmt.Println("请输入要DES解密的文件的路径：")
	lable:
	//获取路径，兼容路径中含有空格
	var src string
	var arr []rune
	var a rune

	for i := 0; ; i++ {
		fmt.Scanf("%c", &a)
		arr = append(arr, a)
		//arr[i] = a err  必须用append增加元素，arr本身长度、容量均为0
		if arr[i] == 10 {
			fmt.Print(string(arr))
			break
		}
	}

	//兼容Windows的.exe和goland  Windows下的.exe 回车键代表\r\n 两个字节 [13 10] goland下回车键代表\n [10]
	if arr[len(arr)-2] == 13 {
		if arr[0]==34 {
			src = string(arr[1:len(arr)-3])
		}else {
			src = string(arr[:len(arr)-2])
		}
	} else {
		//为了在goland中运行 goland中回车键代表\n
		src = string(arr[:len(arr)-1])
	}

	f, err := os.Open(src)
	if err != nil {
		fmt.Println("输入路径有误！请重新输入！")
		goto lable
	}
	defer f.Close()
	fileInfo, err := f.Stat()
	fmt.Println("文件正在进行DES解密，请稍安勿躁……………")
	CbcDecryptFile(src, "./解密/"+fileInfo.Name()[2:])
	fmt.Println("文件DES解密成功！")
}
