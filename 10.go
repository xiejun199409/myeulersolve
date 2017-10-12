package main

import (
	"fmt"
	"time"

)

func main() {
	start := time.Now()
	var primepool [2000000]bool
	primepool[0] = false
	for i := 2 ; i < 2000000; i++{
		primepool[i-1] = true
	}
	for i := 2 ; i * i < 2000000; i++{
		if  primepool[i-1] {
			for j := i*i ; j < 2000000 ; j += i{
				primepool[j-1] = false
			}
		}
	}
	sum := 0
	for i :=2 ; i < 2000000 ;i++{
		if primepool[i-1]{
			sum += i
		}
	}
	fmt.Println("Result is", sum, "\nRunning time is", time.Since(start));
}
