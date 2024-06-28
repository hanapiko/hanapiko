package check

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"ascii/types"
)

// Usage checks if the program was supplied the expected command-line arguments, exits on usage error
func Usage(args []string) types.Data {
	if len(args) == 0 {
		// case: go run .
		os.Exit(0)
	} else if len(args) > 3 {
		// checks if the there are more than three arguements
		PrintUsage()
	}
	var out types.Data

	if len(args) == 1 {
		isValid, _ := Expressions(args[0])
		if isValid {
			PrintUsage()
		} else {
			out.Text = args[0]
		}
	} else if len(args) == 2 {
		isValid, filename := Expressions(args[0])
		if isValid {
			out.Text = args[1]
			out.OutputFile = filename
		} else if !strings.HasPrefix(args[0], "--") {
			out.Text = args[0]
			out.Banner = args[1]
		} else {
			PrintUsage()
		}
	} else if len(args) == 3 {
		isValid, filename := Expressions(args[0])
		if isValid {
			out.OutputFile = filename
			out.Text = args[1]
			out.Banner = args[2]
		} else {
			PrintUsage()
		}
	}
	return out
}

// PrintUsage prints the program usage information, then exits the program with the error code 1.
func PrintUsage() {
	usage := "Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard"
	fmt.Println(usage)
	os.Exit(1)
}

// Text
func Text(text string) string {
	// Replace all the newline characters `\n` in the arguement with `\\n`
	text = strings.ReplaceAll(text, "\n", "\\n")

	if text == "" {
		// Nothing to draw
		os.Exit(0)
	} else if text == "\\n" {
		// if the arguement is "\n" the program prints a new line
		fmt.Println()
		os.Exit(0)
	}

	out := ""
	// checking if the runes of the string in the arguement is not of an ascii decimal value of more than 126
	for _, char := range text {
		if char > '~' {
			fmt.Println("Error : Non Ascii character found!! can not display the graphic representation")
			os.Exit(1)
		} else if char >= ' ' {
			// Ignore ASCII characters before the ' ' (space) character
			out += string(char)
		}
	}

	return out
}

// ArtFile given a banner, returns the name of the file with the graphics for the given banner
func ArtFile(banner string) string {
	switch banner {
	case "standard":
		return "standard.txt"
	case "thinkertoy":
		return "thinkertoy.txt"
	case "shadow":
		return "shadow.txt"
	case "lean":
		return "lean.txt"
	default:
		fmt.Printf("Invalid banner: %q\n", banner)
		PrintUsage()
	}

	return ""
}

// This function is used to check if the flags to be matched are the same.
func Expressions(s string) (bool, string) {
	re := regexp.MustCompile(`^--output=(.+)`)
	if re.MatchString(s) {
		matches := re.FindStringSubmatch(s)
		// fmt.Println(matches)
		if matches[1] == "standard.txt" || matches[1] == "shadow.txt" || matches[1] == "thinkertoy.txt" || matches[1] == "lean.txt" {
			fmt.Println("File already exists, use another file name at the flag")
			os.Exit(1)
		}
		if !strings.HasSuffix(matches[1], ".txt") {
			fmt.Println("Use a txt file for the flag")
			os.Exit(1)
		}
		return true, matches[1]
	}
	return false, ""
}
