package main

import (
	"fmt"
)

func fibo(x int) int{
	if x == 1 {
		return 1
	} else if x== 2 {
		return 2
	} else {
		return fibo(x-1) + fibo(x-2)
	}
}

func main(){
	var sum,j int
	sum = 0
	j = fibo(1)
	for i := 2; j<4000000;i++ {
		if j%2 == 0{
			sum += j
		}
		j = fibo(i)
	}
	fmt.Println(sum)
}

