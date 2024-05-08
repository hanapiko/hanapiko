package myfunctions

func Punctuations(s []string) []string {
	puncs := []string{",", ".", "!", "?", ":", ";"}
	//punc at end
	for i, word := range s {
		for _, punc := range puncs {
			if string(word[0]) == punc && s[len(s)-1] == s[i] {
				s[i-1] = s[i-1] + word
				s = s[:len(s)-1]
			}
		}
	}
	//middle punctuation
	for i, word := range s {
		for _, punc := range puncs {
			if string(word[0]) == punc && string(word[len(word)-1]) == punc && s[i] != s[len(s)-1] {
				s[i-1] = s[i-1] + word
				s = append(s[:i], s[i+1:]...)
			}
		}
	}
	// punc in the middle connect to word
	for i, word := range s {
		for _, punc := range puncs {
			if string(word[0]) == punc && string(word[len(word)-1]) != punc {
				s[i-1] += punc
				s[i] = word[1:]
			}
		}
	}
	 return s
}
