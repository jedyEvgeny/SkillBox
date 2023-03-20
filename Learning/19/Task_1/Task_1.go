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
	arraySecond := [sizeArrTwo]int{32, 64, 65, 666, 39184}
	fmt.Println(newArr(arrayFirst, arraySecond))
}

func newArr(inArrFirst [sizeArrOne]int, inArrSec [sizeArrTwo]int) [sizeArrThree]int {
	var arrayUnity [sizeArrThree]int
	var a, b, i int
	for ; a < sizeArrOne && b < sizeArrTwo; i++ {
		if inArrFirst[a] < inArrSec[b] {
			arrayUnity[i] = inArrFirst[a]
			a++
		} else {
			arrayUnity[i] = inArrSec[b]
			b++
		}
	}
	for ; a < sizeArrOne; i++ {
		arrayUnity[i] = inArrFirst[a]
		a++
	}
	for ; b < sizeArrTwo; i++ {
		arrayUnity[i] = inArrSec[b]
		b++
	}
	return arrayUnity
}
