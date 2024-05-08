package myfunctions

import (
	"strings"
	"strconv"
	"fmt"
)

func Capitalize(s []string) []string {
	for i := 0; i < len(s); i++ {
		if strings.Contains(s[i], "(cap"){
			if strings.Contains(s[i], "(cap,"){
				number, err := strconv.Atoi(s[i+1][:len(s[i+1])-1])
				if err != nil || number > len(s) || number < 0 {
					fmt.Println("Error at conversion  or cap is out of range.", err, number)
					continue
				}
				for j := i - number; j < i;  j++ { // capitalize the elements base on the cap value 
					// if i-number < 0 {
					// 	fmt.Println("error")
					// 	break
					//}
					s[j] = strings.ToUpper(string(s[j][0])) + strings.ToLower(string(s[j][1:]))
				}
				s = append(s[:i], s[i+2:]...)
			} else {
				s[i-1] = strings.ToUpper(string(s[i-1][0])) + strings.ToLower(string(s[i-1][1:]))
				s = append(s[:i], s[i+1:]...)
			}
		}
	}
	return s
}

func Upp(s []string) []string {
	for i := 0; i < len(s); i++ {
		if strings.Contains(s[i], "(up") {
			if strings.Contains(s[i], "(up,") {
				number, err := strconv.Atoi(s[i+1][:len(s[i+1])-1]) // converting the string to integer
				if err != nil {
					fmt.Println(err)
				}
				for j := i - number; j < i; j++ {
					s[j] = strings.ToUpper(s[j])
				}
				s = append(s[:i], s[i+2:]...)
			} else {
				s[i-1] = strings.ToUpper(s[i-1])
				s = append(s[:i], s[i+1:]...)
			}
		}
	}
	return s
}

func Low(s []string) []string {
	for i := 0; i < len(s); i++ {
		if strings.Contains(s[i], "(low") {
			if strings.Contains(s[i], "(low,") {
				number, err := strconv.Atoi(s[i+1][:len(s[i+1])-1])
				if err != nil {
					fmt.Println(err)
				}
				for j := i - number; j < i; j++ {
					s[j] = strings.ToLower(s[j])
				}
				s = append(s[:i], s[i+2:]...)
			} else {
				s[i-1] = strings.ToLower(s[i-1])
				s = append(s[:i], s[i+1:]...)
			}
		}
	}
	return s
}