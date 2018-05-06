package base58

import (
	"fmt"
	"testing"
)

func Test_Encode(t *testing) {
	byteData := []byte{0x00, 0xFF}
	result := Encode(byteData)
	fmt.Printf("编码结果：%v", string(result))
	fmt.Println()
}

func Test_Encode(t *testing) {
	data := "15Q"
	result := Decode(byteData)
	fmt.Print("解码结果：")
	fmt.Println(result)
}
