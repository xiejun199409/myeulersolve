package main

import "fmt"

func main(){
	sum := 0
	for i := 3; i<1000 ;i += 3{
		if i%5 != 0 {
			sum += i
		}
	}
	for j := 5; j<1000 ;j += 5{
		sum += j
	}
	fmt.Println(sum);
}
