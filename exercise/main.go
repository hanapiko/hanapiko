package main

import (
	"fmt"
	"os"
	"reflect"
)

func main() {
	args := os.Args

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	if len(args) != 3 {
		fmt.Println("error usage should be")
	}

	fileContent, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println(err)
	}

	fileContents := string(fileContent)

	expected := `package main
	import "fmt"
	
	func main() {
		fmt.Println("Hello world")
	}`

	if !reflect.DeepEqual(expected, fileContents) {
		fileContents = `package main
		import "fmt"
		
		func main() {
			fmt.Println("Hello world")
		}`
	} else {
		fmt.Print("well formated")
	}
	result := []byte(fileContents)

	os.WriteFile(outputFile, result, 0o644)
}
