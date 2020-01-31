package main

import (
	"fmt"
)

func Contains(a []int, x int) bool {
    for _, n := range a {
        if x == n {
            return true
        }
    }
    return false
}

func FindNext(a []int, x int) int {
	for {
	if !Contains(a, x+1) {
		return x+1
	}
	x = x +1
	}
}

func main() {
	fmt.Println("Hello, playground")
	
	
	var numarr = []int{-10000,1,14,3,2, 1000000}
	
	var small int = 1
	
	for _, k := range numarr {
		if k == small {
			small = FindNext(numarr, small)
		}
	}
	
	fmt.Println(small)
}
