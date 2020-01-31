package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")
	
	myMap := make(map[int]bool) 
	
	var numarr = []int{-10000,1,14,3,2,4,6,5, 1000000}
	
	var small int = 1
	
	for _, k := range numarr {
		myMap[k] = true
		if k == small{
			for {
				if(myMap[small + 1] != true){
					small = small + 1
					break
				} 
				small++
			}
		}
	}
	
	fmt.Println(small)
}
