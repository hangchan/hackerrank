package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the hourglassSum function below.
func hourglassSum(arr [][]int32) int32 {
	//var x int32 = 0
	//var y int32 = 0
	var total int32 = 0
	var maxtotal int32 = -64
	//var totalarr []int32

	/*
		arr[y][x]
		arr[y][x+1]
		arr[y][x+2]
		arr[y+1][x+1]
		arr[y+2][x]
		arr[y+2][x+1]
		arr[y+2][x+1]
	*/

	for j := 0; j < 4; j++ {
		for i := 0; i < 4; i++ {
			total = 0
			for ii := 0; ii < 3; ii++ {
				total += arr[i][ii+j]
				fmt.Println("top", arr[i][ii+j])
			}
			for ii := 0; ii < 1; ii++ {
				total += arr[i+1][ii+1+j]
				fmt.Println("middle", arr[i+1][ii+1+j])
			}
			for ii := 0; ii < 3; ii++ {
				total += arr[i+2][ii+j]
				fmt.Println("bottom", arr[i+2][ii+j])
			}
			fmt.Println(total)
			if total > maxtotal {
				maxtotal = total
			}
		}
	}

	fmt.Println("max:", maxtotal)
	return maxtotal
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	/*stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024 * 1024)*/

	var arr [][]int32
	for i := 0; i < 6; i++ {
		arrRowTemp := strings.Split(readLine(reader), " ")

		var arrRow []int32
		for _, arrRowItem := range arrRowTemp {
			arrItemTemp, err := strconv.ParseInt(arrRowItem, 10, 64)
			checkError(err)
			arrItem := int32(arrItemTemp)
			arrRow = append(arrRow, arrItem)
		}

		if len(arrRow) != int(6) {
			panic("Bad input")
		}

		arr = append(arr, arrRow)
	}

	result := hourglassSum(arr)

	fmt.Printf("%d\n", result)

	//writer.Flush()
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
