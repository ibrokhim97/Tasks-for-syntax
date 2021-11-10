package main

import (
	"fmt"
	"strconv"
)

func main() {
	iNumbertocheck := 91029
	str := strconv.FormatInt(int64(iNumbertocheck), 10)
	b := len(str)

	for i := 0; i < len(str); i++ {
		b = b - 1
		if str[i] != str[b] {
			fmt.Println("Not a Palindrome")
			return
		}
	}
	fmt.Println("it's a Palindrome")
}
