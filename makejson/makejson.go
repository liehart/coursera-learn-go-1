package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	var name, address string
	personMap := make(map[string]string)

	fmt.Printf("Please input your name: ")
	fmt.Scanf("%s", &name)
	fmt.Printf("Please input your address: ")
	fmt.Scanf("%s", &address)

	personMap["name"] = name
	personMap["address"] = address

	dataByte, err := json.Marshal(personMap)
	if err != nil {
		fmt.Printf("Marshalling failed")
		os.Exit(1)
	}
	fmt.Println(string(dataByte))
}
