package myfunctions

// Vowel function takes a slice of strings as input and checks for vowels to determine if "a" or "an" should be used.
func Vowel(s []string) []string {
	vowels := []string{"a", "e", "i", "o", "u", "h", "A", "E", "I", "O", "U", "H"}
	for i := 0; i < len(s); i++ {
		for _, fletter := range vowels {
			if s[i] == "a" && string(s[i+1][0]) == fletter {
				s[i] = "an"
			} else {
				if s[i] == "A" && string(s[i+1][0]) == fletter {
					s[i] = "An"
				}
			}
		}
	}
	return s
}
