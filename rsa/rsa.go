package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
)

// 生成rsa的密钥对, 并且保存到磁盘文件中
func GenerateKey(keySize int) {
	// 1. 使用rsa中的GenerateKey方法生成私钥
	privateKey, err := rsa.GenerateKey(rand.Reader, keySize)
	if err != nil {
		panic(err)
	}
	// 2. 通过x509标准将得到的ras私钥序列化为ASN.1 的 DER编码字符串
	derText := x509.MarshalPKCS1PrivateKey(privateKey)
	// 3. 要组织一个pem.Block
	block := pem.Block{
		Type : "rsa private key", // 这个地方写个字符串就行
		Bytes : derText,
	}
	// 4. pem编码
	file, err := os.Create("private.pem")
	if err != nil {
		panic(err)
	}
	pem.Encode(file, &block)
	file.Close()

	// ============ 公钥 ==========
	// 1. 从私钥中取出公钥
	publicKey := privateKey.PublicKey
	// 2. 使用x509标准序列化
	derstream, err := x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		panic(err)
	}
	// 3. 将得到的数据放到pem.Block中
	block = pem.Block{
		Type : "rsa public key",
		Bytes : derstream,
	}
	// 4. pem编码
	file, err  = os.Create("public.pem")
	if err != nil {
		panic(err)
	}
	pem.Encode(file, &block)
	file.Close()
}

//RSA加密，公钥加密
func Encrypt(plainText []byte,filePath string)[]byte  {
	// 1. 打开文件, 并且读出文件内容
	f,err:=os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	fileInfo,err:=f.Stat()
	if err != nil {
		panic(err)
	}
	buf:=make([]byte,fileInfo.Size())
	f.Read(buf)
	// 2. pem解码
	block,_:=pem.Decode(buf)
	pubInterface,err:=x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	//pubkey,_:=pubInterface.(*rsa.PublicKey)
	pubKey,_:=pubInterface.(*rsa.PublicKey)
	// 3. 使用公钥加密
	cipherText,err:=rsa.EncryptPKCS1v15(rand.Reader,pubKey,plainText)
	if err != nil {
		panic(err)
	}
	return cipherText

}

// RSA 解密,私钥解密
func Decrypt(cipherText []byte,filePath string)[]byte  {
	// 1. 打开文件, 并且读出文件内容
	f,err:=os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	fileInfo,err:=f.Stat()
	if err != nil {
		panic(err)
	}
	buf:=make([]byte,fileInfo.Size())
	f.Read(buf)
	// 2. pem解码
	block,_:=pem.Decode(buf)
	privKey,err:=x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	// 3. 使用私钥解密
	plainText,err:=rsa.DecryptPKCS1v15(rand.Reader,privKey,cipherText)
	if err != nil {
		panic(err)
	}
	return plainText

}


