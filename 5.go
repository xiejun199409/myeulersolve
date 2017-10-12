package main

import (
	"fmt"
	"time"
)

func main(){
	start := time.Now()
	var result int
	for result = 1 ;;result++{
		for i := 2 ; i <=20 ; i++{
			if result%i != 0 {
				break
			}
			if i == 20 {
				fmt.Println("Result is",result,"\nRunning time is",time.Since(start));
				return
			}
		}
	}
}
