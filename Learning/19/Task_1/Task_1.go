package main

import (
	"fmt"
)

const (
	sizeArrOne   = 4
	sizeArrTwo   = 5
	sizeArrThree = sizeArrOne + sizeArrTwo
)

func main() {
	arrayFirst := [sizeArrOne]int{10, 16, 32, 64}
	arraySecond := [sizeArrTwo]int{32, 64, 15, 258, 139}
	fmt.Println(newArr(arrayFirst, arraySecond))
}

func newArr(inArrFirst [sizeArrOne]int, inArrSec [sizeArrTwo]int) [sizeArrThree]int {
	var arrayUnity [sizeArrThree]int
	for i := 0; i < sizeArrOne; i++ {
		arrayUnity[i] = inArrFirst[i]
	}
	for i := sizeArrOne; i < sizeArrThree; i++ {
		arrayUnity[i] = inArrSec[i-sizeArrOne]
	}
	return arrayUnity
}
