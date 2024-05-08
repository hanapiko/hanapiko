package banner

import (
	"log"
	"os"
	"strings"
)

func Art(bannerFile string) map[rune][]string {
	text, err := os.ReadFile(bannerFile)
	if err != nil {
		log.Fatalf("Error reading file: \"%s\"\n", bannerFile)
	}
	
	s := string(text)
	data := make(map[rune][]string)
	lines := strings.Split(s, "\n")

	r := ' '
	var g [] string

	for i, line := range lines {
		if i == 0 {
			g = make([]string, 0)
		} else if i % 9 == 0 {
			data[r] = g
			g = make([]string, 0)
			r++
		} else {
			g = append(g, line)
		}
	}
	data[r] = g

	return data
}