package main

import (
	"fmt"
	"os"
	"strings"

	"ascii/check"
	"ascii/file"
	"ascii/graphics"
)

func main() {
	args := os.Args[1:]
	data := check.Usage(args)

	text := check.Text(data.Text)

	banner := "standard"
	if data.Banner != "" {
		banner = data.Banner
	}
	asciiArtFile := check.ArtFile(banner)
	bannerFileContents := file.ReadArtFile(asciiArtFile)

	// the variable inputArgs splits the string(arguement) into substrings, by a separator "\n"
	inputArgs := strings.Split(text, "\\n")

	output := graphics.Graphic(inputArgs, bannerFileContents)
	if data.OutputFile != "" {
		// wrte output to data.OutputFile
		result := []byte(output)
		err := os.WriteFile(data.OutputFile, result, 0644)
		if err != nil {
			fmt.Println("Error writing to file", err)
			os.Exit(1)
		}
		return
	}
	fmt.Print(output)
}
