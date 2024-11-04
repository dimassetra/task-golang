package main

import (
	"Golang/utils"
	"fmt"
)

func main() {
	n := 3

	fmt.Println("Welcome")
	fmt.Println("Nilai MagicSum: ", utils.MagicSum(n))
	fmt.Println("Nilai MagicPow: ", utils.MagicPow(n))
	fmt.Println("Nilai MagicOdd: ", utils.MagicOdd(n))
	fmt.Println("Nilai MagicGrade: ", utils.MagicGrade(n))
	fmt.Println("Nilai MagicName: ", utils.MagicName(n))
	fmt.Println("Nilai MagicTria: ", utils.MagicTria(n))

}
