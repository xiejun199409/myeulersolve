package main

import (
	"fmt"
	"math"

	"time"
)

func main(){
	var num,i int64
	start := time.Now()
	num = 600851475143
	poolsize := int64(math.Sqrt(float64(num)))
	for i = 2 ; i <= poolsize ; i++{
		for true {
			if (num%i == 0){
				num /= i
			}else{
				if num == 1{
					num = i
				}
				break
			}
		}
	}
	fmt.Println("Result is",num,"\nRunning time is",time.Since(start));

}
