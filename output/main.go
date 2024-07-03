package main

import (
	"fmt"
	"os"
	"strings"

	"asciiartfs/functions"
)

func main() {
	usage := "Usage: go run . [STRING] [BANNER]\n\nEX: go run . something standard"
	if len(os.Args) == 1 {
		fmt.Println(usage)
		return
	} else if len(os.Args) > 3 {
		// checks if the there are more than three arguements.
		fmt.Println(usage)
		return
	}

	args := os.Args[1]

	// replaces all the new lines in the arguement with a new line
	args = strings.ReplaceAll(args, "\n", "\\n")

	if args == "\\n" {
		// if  the arguement is "\n" the program prints a new line
		fmt.Println()
		return
	}

	if args == "" {
		return
	}
	// checking if the runes of the string in the arguement is not of an ascii decimal value of more than 126
	for _, chr := range args {
		if chr > '~' {
			fmt.Println("Error : Non Ascii character found!! can not display the graphic representation")
			return
		}
	}

	// asigning a variable, asciiArtFile that's going to store the value of the banner files
	asciiArtFile := "standard"
	if len(os.Args) > 2 {
		asciiArtFile = os.Args[2]
	}

	switch asciiArtFile {
	case "standard":
		asciiArtFile = "standard.txt"
	case "thinkertoy":
		asciiArtFile = "thinkertoy.txt"
	case "shadow":
		asciiArtFile = "shadow.txt"
	case "lean":
		asciiArtFile = "lean.txt"
	default:
		fmt.Println(usage)
		return
	}

	// the variable inputArgs splits the string(arguement) into substrings, by a separator "\n"
	inputArgs := strings.Split(args, "\\n")
	// read the bannerfile
	asciiArt, err := os.ReadFile(asciiArtFile)
	if err != nil {
		fmt.Printf("Error reading file: %q\n%v\n", asciiArtFile, err)
		return
	}

	if len(asciiArt) == 0 {
		fmt.Println("Error: Empty banner file")
		return
	}

	// the variable asciiArtFile stores the data read from bytes to string
	asciiArtString := string(asciiArt)

	var bannerFileContents []string
	if asciiArtFile == "thinkertoy.txt" {
		bannerFileContents = strings.Split(asciiArtString, "\r\n")
	} else {
		bannerFileContents = strings.Split(asciiArtString, "\n")
	}

	output := functions.Graphic(inputArgs, bannerFileContents)
	fmt.Print(output)
}
