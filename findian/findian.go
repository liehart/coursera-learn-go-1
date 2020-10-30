package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Printf("Enter a string: ")
	inputReader := bufio.NewReader(os.Stdin)
	Text, _ := inputReader.ReadString('\n')
	Text = strings.ToLower(Text)
	Text = strings.Replace(Text, " ", "", -1)
	TextLength := len(Text) - 1
	if TextLength >= 3 {
		TextFirst := Text[0] == 'i'
		TextLast := Text[TextLength-1] == 'n'
		IsValid := TextFirst && TextLast
		for i := 1; i < TextLength-1; i++ {
			if Text[i] == 'a' {
				if IsValid {
					fmt.Println("Found!")
					return
				}
			}
		}
	}
	fmt.Println("Not Found!")
}
