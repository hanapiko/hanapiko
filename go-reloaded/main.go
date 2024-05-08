package main

import (
	"fmt"
	"go-reloaded/myfunctions"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 3 { // checking if exactly 3 arguments are provided
		fmt.Println("Expecting 3 arguments")
		return
	}
	args := os.Args[1:]
	text, err := os.ReadFile(args[0]) // reading the file of the content at the specified argument
	if err != nil {
		fmt.Println(err)
	}
	// a := strings.Split(string(text), "\n")
	// empty := ""

	// for _, ch := range a {

		//fmt.Println(string(text))
		words := strings.Fields(string(text))
		words = myfunctions.Capitalize(words)
		words = myfunctions.Upp(words)
		words = myfunctions.Low(words)
		words = myfunctions.Hex(words)
		words = myfunctions.Bin(words)
		words = myfunctions.Vowel(words)
		words = myfunctions.Punctuations(words)
		words = myfunctions.Apostrophe(words)
		
		//empty += result
	
	result := strings.Join(words, " ") + "\n"
	results := []byte(result)
	os.WriteFile(args[1], results, 0644)
}
