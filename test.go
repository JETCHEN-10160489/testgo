package main

import (
	"fmt"
)
//test123
func main() {
	for i:=1;i<=3;i++ {
		if i>=2 {
			fmt.Print("当前的i值为：")
			fmt.Print(i)
			fmt.Print("\n")
			continue
		}
		fmt.Print("testgo\n")
	}

}
