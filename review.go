package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	scanner := bufio.NewScanner(os.Stdin)

	var x int
	var y string

	fmt.Scanf("%d", &x)
	for i := 0; i < x; i++ {
		scanner.Scan()
		y = scanner.Text()
		//fmt.Println(y)
		z := strings.Split(y, "")
		//fmt.Println(z)
		oddz := []string{}
		evenz := []string{}
		for i, v := range z {
			if i == 0 || i%2 == 0 {
				oddz = append(oddz, v)
			} else {
				evenz = append(evenz, v)
			}
		}
		for _, v := range oddz {
			fmt.Printf("%s", v)
		}
		fmt.Printf(" ")
		for _, v := range evenz {
			fmt.Printf("%s", v)
		}
		fmt.Printf("\n")
	}
}
