package main

import (
	"fmt"
	"time"
	"os"
)

func main() {
	start := time.Now()
	file,err := os.Open("11in.txt")
	if err != nil{
		fmt.Println("File reading error")
		return
	}
	defer file.Close()
	stat,err := file.Stat()
	if err != nil{
		fmt.Println("Unable to get size")
		return
	}
	buf := make([]byte, stat.Size())
	_,err =file.Read(buf)
	if err != nil{
		fmt.Println("Unable to read file")
	}
	area := make([]int , 400)
	count := 0
	for i := 0 ; i < int(stat.Size()) ; i++{
		if (buf[i] <=  '9') && (buf[i] => '0') && (buf[i+1] <= '9') && (buf[i+1] => '0') && buf[i+2] > '9' && buf[i+2] < '0'{

		}
	}
	fmt.Println("Result is", num, "\nRunning time is", time.Since(start));
}
