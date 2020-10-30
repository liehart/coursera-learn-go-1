package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var intSlice = make([]int, 3)
	var input string
	for {
		fmt.Printf("Please input an integer number: ")
		fmt.Scanf("%s", &input)
		if strings.ToLower(input) == "x" {
			fmt.Println("Program ended")
			os.Exit(0)
		}
		intVal, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Input invalid, try again")
			continue
		}
		intSlice = append(intSlice, intVal)
		sort.Sort(sort.IntSlice(intSlice))
		fmt.Println(intSlice[3:len(intSlice)])
	}
}
