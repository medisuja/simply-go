package main

import (
	"fmt"
	"strings"
)

func main() {
	check := validOrNotValid("00001111100")
	fmt.Println(check)
}

func validOrNotValid(s string) string {
	chars := strings.Split(s, "")
	
	first := 0
	last := 0
	
	fmt.Println("LEFT PAD")
	// left pad (take the index whose value is 1 from the left)  
	for first = 0; first < len(s); first++ {
		fmt.Println(first, chars[first])
		if chars[first] == "1" {
			break
		}
	}
	
	fmt.Println("=======================")

	fmt.Println("RIGHT PAD")
	// right pad (take the index whose value is 1 from the right)
	for last = len(s) - 1; last >= 0; last-- {
		fmt.Println(last, chars[last])
		if chars[last] == "1" {
			break
		}
	}

	fmt.Println("=======================")
	
	fmt.Println(first, "(VALUE INDEX FROM LEFT PAD)")
	
	fmt.Println(last, "(VALUE INDEX FROM RIGHT PAD)")

	fmt.Println("=======================")
	
	for i := first + 1; i < last; i++ {
		if chars[i] == "0" {
			return "NOT VALID";
		}
	}

	return "VALID";
}
