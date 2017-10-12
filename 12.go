package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	for i := 1 ; ; i++{
		sum := (1+i)*i/2
		count := 0
		var j int
		for j = 1; j * j < sum ; j++{
			if sum%j == 0 {
				count++
			}
		}
		if j * j == sum{
			count = count * 2 + 1
		}else{
			count *= 2
		}
		if count > 500 {
			fmt.Println("Result is", sum, "\nRunning time is", time.Since(start))
			return
		}
	}
}