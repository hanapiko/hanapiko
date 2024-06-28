package file

import (
	"fmt"
	"os"
	"regexp"
)

// ReadArtFile reads all the lines in the given ascii art file, then returns all the lines read
func ReadArtFile(asciiArtFile string) []string {
	// read the bannerfile
	asciiArt, err := os.ReadFile(asciiArtFile)

	if len(asciiArt) == 0 {
		fmt.Println("Empty banner file")
		os.Exit(1)
	}

	if err != nil {
		fmt.Printf("Error reading file: %q\n%v\n", asciiArtFile, err)
		os.Exit(1)
	}

	// the variable asciiArtFile stores the data read from bytes to string
	asciiArtString := string(asciiArt)

	re := regexp.MustCompile(`\r?\n`)
	bannerFileContents := re.Split(asciiArtString, -1)

	return bannerFileContents
}
