package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var _ = strconv.Itoa // Ignore this comment. You can still use the package "strconv".

	var i uint64 = 4
	var d float64 = 4.0
	var s string = "HackerRank "

	scanner := bufio.NewScanner(os.Stdin)
	// Declare second integer, double, and String variables.
	var x uint64
	var y float64
	var z string
	// Read and save an integer, double, and String to your variables.
	// fmt.Scanf("%d", &x)
	scanner.Scan()
	fmt.Scanf("%f", &x)
	fmt.Scanf("%f", &y)
	z = scanner.Text()
	// Print the sum of both integer variables on a new line.
	fmt.Println(i + x)
	// Print the sum of the double variables on a new line.
	fmt.Println(d + y)
	// Concatenate and print the String variables on a new line
	fmt.Println(s + z)
	// The 's' variable above should be printed first.
}