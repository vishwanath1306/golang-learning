package main

import "fmt"

const (
	//Big is used to represent a very large number
	Big = 1 << 100
	//Small is uesd to represent a very small number
	Small = Big >> 99
)

func needInt(x int) int {
	return x*10 + 1
}
func needFloat(x float64) float64 {
	return x * 0.1
}
func main() {

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
