package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	var primepool [10001]int
	count := 1
	primepool[0] = 2
	for i := 3 ; count < 10001 ; i++{
		for j :=0 ; j<count;j++{
			if i%primepool[j] == 0{
				break
			}
			if j == count - 1{
				primepool[count] = i
				count++
			}
		}
	}
	fmt.Println("Result is", primepool[10000], "\nRunning time is", time.Since(start));
}
