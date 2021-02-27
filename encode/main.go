package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var conv string = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

func main() {
	str := os.Args[1]
	bitArray := toBitArray(toBinary(str))
	encoded := toBase64Str(bitArray)
	fmt.Println(encoded)
}

func toBinary(str string) []byte {
	return []byte(str)
}

func toBitArray(byteArray []byte) []uint64 {
	var bitArray []uint64
	for i := 0; i < len(byteArray); i++ {
		bt := byteArray[i]
		binStr := fmt.Sprintf("%08b\n", bt)
		intArray := stringToIntArray(binStr)
		bitArray = append(bitArray, intArray...)
	}
	return bitArray
}

func stringToIntArray(str string) []uint64 {
	var intArray []uint64
	for i := 0; i < len(str)-1; i++ {
		s := string(str[i])
		n, err := strconv.Atoi(s)
		if err != nil {
			panic("stringToIntArray error")
		}
		intArray = append(intArray, uint64(n))
	}

	return intArray
}

func toBase64Str(bitArray []uint64) string {
	r := len(bitArray) % 6
	if r != 0 {
		toAdd := make([]uint64, 6-r)
		bitArray = append(bitArray, toAdd...)
	}

	var base64StrArray []string
	for i := 0; i < len(bitArray); i += 6 {
		sliced := bitArray[i : i+6]
		b := binToBase64Str(fmt.Sprintf("%d%d%d%d%d%d", sliced[0], sliced[1], sliced[2], sliced[3], sliced[4], sliced[5]))
		base64StrArray = append(base64StrArray, string(b))
	}

	r = len(base64StrArray) % 4
	if r != 0 {
		for i := 0; i < r; i++ {
			base64StrArray = append(base64StrArray, "=")
		}
	}

	return strings.Join(base64StrArray, "")
}

func binToBase64Str(str string) byte {
	i, err := strconv.ParseInt(str, 2, 64)
	if err != nil {
		panic("binToBase64Str error")
	}

	s := conv[i]
	return s
}
