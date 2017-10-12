package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	var squareSum,sumSquare,sum,result int
	squareSum = 0
	sum = 0
	for i := 1 ; i <= 100 ; i++{
		squareSum += i*i
	}
	for j := 1 ; j <= 100 ; j++{
		sum += j
	}
	sumSquare = sum * sum
	if sumSquare > squareSum {
		result = sumSquare - squareSum
	}else{
		result = squareSum - sumSquare
	}
	fmt.Println("Result is", result, "\nRunning time is", time.Since(start));
}