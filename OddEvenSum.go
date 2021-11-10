package main

import "fmt"

func main() {
	fmt.Println("Enter Number :")
	var n int
	fmt.Scanln(&n)

	if n&2 == 0 {
		fmt.Println(n, "Is even number")
	} else {
		fmt.Println(n, "is a odd number")
	}
}
