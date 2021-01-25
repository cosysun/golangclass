package main

import "fmt"

func main() {
	var bByte byte = 1
	var sByte byte = 's'
	fmt.Printf("bByte:%d, sByte:%d\n", bByte, sByte)

	str := "Golang"
	fmt.Println(str)
	fmt.Println(StringToByte(str))
	fmt.Println(ByteToString(StringToByte(str)))
}

func ByteToString(bstr []byte) string {
	return string(bstr)
}

func StringToByte(str string) []byte {
	return []byte(str)
}
