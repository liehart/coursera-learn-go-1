package main

import "fmt"

func main() {
	var FloatNumber float32
	fmt.Printf("Please input a floating number: ")
	fmt.Scan(&FloatNumber)
	fmt.Println("Truncated number is", int32(FloatNumber))
}
