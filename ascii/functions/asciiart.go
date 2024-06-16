package functions

func Dev(splitargs, char []string) string {
	a := 0
	b := ""
	for _, word := range splitargs {
		if word == "" {
			a = 1
		} else {
			a = 8
		}
		for i := 1; i <= a; i++ {
			for _, runWord := range word {
				index := (runWord-32)*9 + rune(i)
				b += char[index]
			}
			b += "\n"
		}
	}
	return b
}