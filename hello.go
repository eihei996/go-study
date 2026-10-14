package main

import (
	"fmt"

	"demo/mathtool"
)

func main() {
	var x, y int
	fmt.Print("请输入两个整数：")
	fmt.Scan(&x, &y)
	fmt.Println("hello,go go go?")
	fmt.Println(mathtool.Add(x, y))
	fmt.Println(mathtool.Sub(x, y))
}
