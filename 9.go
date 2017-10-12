package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	for i := 1 ; i < 33 ; i++{
		for j := i+1 ; i + 2 * j < 100 ; j++{
			k := 100 - i - j
			if i * i + j * j == k * k
		}
	}
	fmt.Println("Result is", num, "\nRunning time is", time.Since(start));
}
