package main

import "fmt"

func print(num int, text string) {
	fmt.Printf("%v is %v \n", num, text)
}

func even(num int) bool {
	if num%2 == 0 {
		return true
	} else {
		return false
	}
}

func odd(num int) bool {
	if num%2 == 1 {
		return true
	} else {
		return false
	}
}

func iterateOverNums(nums []int) {
	for _, num := range nums {
		if even(num) {
			print(num, "even")
		} else if odd(num) {
			print(num, "odd")
		}
	}
}
