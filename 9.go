package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	for i := 1 ; i < 333 ; i++{
		for j := i+1 ; i + 2 * j < 1000 ; j++{
			k := 1000 - i - j
			if (i * i + j * j) == k * k {
				fmt.Println("Result is", i * j * k , "\nRunning time is", time.Since(start));
				return
			}
		}
	}
}
