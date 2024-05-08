package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	a := ""
	for scanner.Scan() {
		a += scanner.Text() + "\n"
	}
	if strings.Contains(a, "package main\nimport \"fmt\"\n\nfunc main() {\n   fmt.Println(\"hello world\")\n}") {
		write := os.WriteFile("output.txt", []byte("well formatted"), 0644)
		if write != nil {
			fmt.Println("error writing file")
		}
	} else {
		apple:= "package main\nimport \"fmt\"\n\nfunc main() {\n   fmt.Println(\"hello world\")\n}"
		wri := os.WriteFile("output.txt",[]byte(apple), 0644)
		if wri != nil {
			fmt.Println("error writing file")
		}
	}
	fmt.Println("well formatted")

}
