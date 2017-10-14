package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	bigSeq := 0
	bigNum := 0
	for i := 1 ; i < 1000000 ; i++{
		tempNum := i
		tempSeq := 0
		for tempNum != 4{
			if tempNum % 2 == 0{
				tempNum /= 2
				tempSeq++
			}else{
				tempNum = 3*tempNum +1
				tempSeq++
			}
		}
		if tempSeq > bigSeq {
			bigSeq = tempSeq
			bigNum = i
		}
	}
	fmt.Println("Result is", bigNum, "\nRunning time is", time.Since(start));
}
