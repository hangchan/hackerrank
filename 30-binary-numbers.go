package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func convertBinary(n int32) {
	b := strconv.FormatInt(int64(n), 2)
	// fmt.Println(b)
	var count int32
	var maxCount int32

	for _, v := range b {
		if v == 48 {
			count = 0
			//fmt.Println("I'm 0")
		} else {
			//fmt.Println(v)
			//fmt.Println(rune(0))
			count++
			if count > maxCount {
				maxCount = count
			}
		}
	}

	//fmt.Println(b)
	fmt.Println(maxCount)

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)
	convertBinary(n)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
