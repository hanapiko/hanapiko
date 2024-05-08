package main

import (
	"fmt"
	"os"

	//"strings"

	//"log"
	"asciiart/print"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("No arguments found")
		os.Exit(0)
	}

	art := print.ArtWork(args[0], "standard.txt")
	fmt.Print(art)
}

// text, err := os.ReadFile("standard.txt")
//  if err != nil {
//  	log.Fatalln("Error reading file", err)
//  }
// fmt.Println(string(text))

// if len(args) != 1 {
// fmt.Println("An error occured")
// }

// ascii := string("standard.txt")
// if ascii == "" {
// 	return
// }
// ascii, err := os.ReadFile("standard.txt")
// if err != nil {
// 	fmt.Println("Error reading file")
// }
// fmt.Println(string(ascii))
// fmt.Println(len(ascii))

// lines := strings.Split(string(ascii), "\n")

//  for i, l := range lines {
//  	fmt.Println(i,l)
//  }
// fmt.Println(len(lines))
// }
