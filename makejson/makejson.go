package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	personMap := make(map[string]string)

	in := bufio.NewReader(os.Stdin)

	fmt.Printf("Please input your name: ")
	name, _ := in.ReadString('\n')
	fmt.Printf("Please input your address: ")
	address, _ := in.ReadString('\n')

	personMap["name"] = name[:len(name)-1]
	personMap["address"] = address[:len(address)-1]

	dataByte, err := json.Marshal(personMap)
	if err != nil {
		fmt.Printf("Marshalling failed")
		os.Exit(1)
	}
	fmt.Println(string(dataByte))
}
