package main

import (
	asciirev "ascii/functions"
	"fmt"
	"os"
	//"regexp"
	"strings"
)

func main() {
	
	if len(os.Args) < 2{
		//fmt.Println("no arguments found")
		return
	}
	args := os.Args[1]
	if args == "" {
		return
	}
	if args == "\\n" {
		fmt.Println()
		return
	}

	for _, chr := range args {
		if chr > '~' {
			fmt.Println("Error : Non Ascii character found!! can not display the graphic representation")
			return
		}
	}
	//re := regexp.MustCompile("\\n")
	args = strings.ReplaceAll(args, "\n", "\\n")
	splitargs := strings.Split(args, "\\n")
	
	ascii, err := os.ReadFile("standard.txt")	
	if err != nil {
		fmt.Println(err)
	}
	asciistring := string(ascii)
	char := strings.Split(asciistring, "\n")
	
	output := asciirev.Dev(splitargs, char)
	fmt.Print(output)
}