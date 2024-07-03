package functions

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestGraphic(t *testing.T) {
	tests := []struct {
		name     string
		str      string
		expected []string
	}{
		{"Single letter", "H", artH},
		{"Hello art", "Hello", artHello},
	}

	asciiArtFile := "../standard.txt"

	// reading the bannerfile
	asciiArt, err := os.ReadFile(asciiArtFile)
	if err != nil {
		fmt.Println("Error reading file", err)
		t.FailNow()
	}

	asciiArtString := string(asciiArt)

	var bannerFileContents []string
	if asciiArtFile == "thinkertoy.txt" {
		bannerFileContents = strings.Split(asciiArtString, "\r\n")
	} else {
		bannerFileContents = strings.Split(asciiArtString, "\n")
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			str := []string{tt.str}
			output := Graphic(str, bannerFileContents)
			expected := artString(tt.expected)
			if output != expected {
				t.Errorf("got: \n%v\nexpected: \n%v\n", output, expected)
			}
		})
	}
}

func artString(art []string) string {
	return strings.Join(art, "\n") + "\n"
}

var artH = []string{
	" _    _  ",
	"| |  | | ",
	"| |__| | ",
	"|  __  | ",
	"| |  | | ",
	"|_|  |_| ",
	"         ",
	"         ",
}

var artHello = []string{
	" _    _          _   _          ",
	"| |  | |        | | | |         ",
	"| |__| |   ___  | | | |   ___   ",
	`|  __  |  / _ \ | | | |  / _ \  `,
	"| |  | | |  __/ | | | | | (_) | ",
	`|_|  |_|  \___| |_| |_|  \___/  `,
	"                                ",
	"                                ",
}
