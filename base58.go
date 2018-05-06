package base58

import (
	"bytes"
	"math/big"
)

const (
	encodebtc = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
)

// 编码思路，将字符串转为字节数组，字节数组转为大数字除以58，取余数添加到结果，商做继续除法操作
func Encode(byteData []byte) []byte {
	//字节切换大数字，除数;Abs 防负数
	bigInt := new(big.Int).SetBytes(byteData)
	bigInt.Abs(bigInt)
	//58,被除数
	big58 := big.NewInt(58)
	//声明变量商，这里我们需要确认下商的大小
	//编码的本质是256转58，256是8位，58是5.857980995位，大约扩容8/5.857980995=1.365658 倍
	//为了安全起见，选1.38倍

	resultLen := len(byteData) * 138 / 100

	resultByte := make([]byte, resultLen+1)

	//循环
	tmp := new(big.Int)

	for bigInt.Sign() > 0 {
		bigInt, tmp = bigInt.DivMod(bigInt, big58, tmp)
		//逆序排列
		resultByte[resultLen] = encodebtc[int32(tmp.Int64())]
		resultLen--
	}

	//首段字节0转“1”
	for i := 0; i < len(byteData); i++ {
		if byteData[i] != 0 {
			break
		}
		resultByte[resultLen] = '1'
		resultLen--
	}
	//未用位置舍弃
	return resultByte[resultLen+1:]
}

//解码的思路和编码相反
func Decode(byteData []byte) []byte {
	resultBigInt := big.NewInt(0)

	for _, value := range byteData {
		index := bytes.IndexByte([]byte(encodebtc), value)
		resultBigInt.Mul(resultBigInt, big.NewInt(58))
		resultBigInt.Add(resultBigInt, big.NewInt(int64(index)))
	}

	result := resultBigInt.Bytes()
	for i := 0; i < len(byteData); i++ {
		if byteData[i] != encodebtc[0] {
			break
		}
		result = append([]byte{0x00}, result...)
	}

	return result
}
