package print

import (
	"asciiart/banner"
	"strings"
)

func ArtWork(s, bannerFile string) string {
	if s == "" {
		return ""
	}

	asciiMap := banner.Art(bannerFile)

	var b strings.Builder

	for i := 0; i < 8; i++ {
		for _, g := range s {
			g := asciiMap[g]
			b.WriteString(g[i])
		}
		b.WriteRune('\n')
	}
	return b.String()
}