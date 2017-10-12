package main

import (
	"fmt"
	"time"
)

func numcheck(x int) bool{
	str := fmt.Sprintf("%d",x)
	num := []byte(str)
	for i := 0; i < len(num) ; i++ {
		if num[i] != num[len(num)-i-1]{
			return false
		}
	}
	return true
}

func main(){
	start := time.Now()
	var num1,num2,sum int
	last := 0
	for num1 = 999 ; num1 > 99 ; num1--{
		for num2 = 999;	num2 > 99 ; num2-- {
			sum = num1 * num2
			if numcheck(sum) == true{
				if last < sum {
					last = sum
				}
			}

		}
	}
	fmt.Println("Result is",last,"\nRunning time is",time.Since(start));
}