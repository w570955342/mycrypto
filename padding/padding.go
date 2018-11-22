package padding

import (
	"bytes"
)

//为最后分组填充
func PaddingLastGroup(plainText []byte, blockSize int) []byte {
	padNum := blockSize - len(plainText)%blockSize
	slice := []byte{byte(padNum)}
	padSlice := bytes.Repeat(slice, padNum)
	newPlainText := append(plainText, padSlice...) //尽量做到不改变传进来的参数
	return newPlainText
}

//删除最后分组填充
func UnPaddingLastGroup(plainText []byte) []byte {
	lenth:=len(plainText)
	padNum:=int(plainText[lenth-1])
	newPlainText:=plainText[:lenth-padNum]
	return newPlainText
}
