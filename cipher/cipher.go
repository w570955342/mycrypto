package cipher

import (
	"crypto/des"
	"mycrypto/padding"
	"crypto/cipher"
	"crypto/aes"
)

//DES加密 分组模式CBC
func DesCbcEncrypt(plainText,key []byte) []byte {
	block,err:=des.NewCipher(key)
	if err!=nil	{
		panic(err)
	}
	padPlainText:=padding.PaddingLastGroup(plainText,block.BlockSize())
	iv:=[]byte("12345678")
	blockMode:=cipher.NewCBCEncrypter(block,iv)
	cipherText:=make([]byte,len(padPlainText))
	blockMode.CryptBlocks(cipherText,padPlainText)

	return cipherText
}

//DES解密 分组模式CBC
func DesCbcDecrypt(cipherText, key []byte) []byte {
	// 1. 建一个底层使用des的密码接口
	block, err := des.NewCipher(key)
	if err != nil {
		panic(err)
	}
	// 2. 创建一个使用cbc模式解密的接口
	iv := []byte("12345678")
	blockMode := cipher.NewCBCDecrypter(block, iv)
	// 3. 解密
	padPlainText:=make([]byte,len(cipherText))
	//fmt.Printf("%d\n", len(cipherText))
	blockMode.CryptBlocks(padPlainText, cipherText)
	// 4. plainText现在存储的是明文, 需要删除加密时候填充的尾部数据
	plainText := padding.UnPaddingLastGroup(padPlainText)
	return plainText
}

// AES加密, 分组模式CTR
func AesCtrEncrypt(plainText, key []byte) []byte {
	// 1. 建一个底层使用AES的密码接口
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	// 2. 创建一个使用CTR分组接口
	iv := []byte("12345678abcdefgh")
	stream := cipher.NewCTR(block, iv)
	// 4. 加密
	cipherText := make([]byte, len(plainText))
	stream.XORKeyStream(cipherText, plainText)

	return cipherText
}

// AES解密, 分组模式CTR
func AesCtrDecrypt(cipherText, key []byte) []byte {
	// 1. 建一个底层使用AES的密码接口
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	// 2. 创建一个使用CTR模式解密的接口
	iv := []byte("12345678abcdefgh")
	stream := cipher.NewCTR(block, iv)
	// 3. 解密
	plainText := make([]byte, len(cipherText))
	stream.XORKeyStream(plainText, cipherText)

	return plainText
}
