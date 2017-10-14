//My system is windows. So the every number ends in crlf in my input. Take care if you run my program in *nix systems

package main

import (
	"fmt"
	"time"
	"math/big"
	"os"
)

func main() {
	start := time.Now()
	sum := big.NewInt(0)
	file,_ := os.Open("13.in")
	defer file.Close()
	stat,_ := file.Stat()
	bs := make([]byte, stat.Size())
	file.Read(bs)

	for i := 0 ; i < 100 ; i++{
		temp := bs[52*i:52*i+50]
		tempString := string(temp)
		tempNum := big.NewInt(0)
		tempNum.SetString(tempString,10)
		fmt.Println(tempNum.String())
		sum.Add(sum,tempNum)
	}
	fmt.Println("Result is", sum.String(), "\nRunning time is", time.Since(start));
}
