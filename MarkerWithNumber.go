package reloaded

import (
	"strconv"
	"strings"
)

func MarkerWithNumber(input string) string {
	words := strings.Fields(input)
	for i := 0; i < len(words); i++ {
		if strings.HasPrefix(words[i], "(low,") || strings.HasPrefix(words[i], "(up,") || strings.HasPrefix(words[i], "(cap,") {
			parts := strings.Split(words[i], ",")
			if len(parts) == 2 {
				number, err := strconv.Atoi(words[i+1][:len(words[i+1])-1])
				words[i+1] = ""
				if err == nil && i-number >= 0 {
					if strings.HasPrefix(words[i], "(low,") {
						for j := i - number; j < i; j++ {
							words[j] = strings.ToLower(words[j])
						}
					} else if strings.HasPrefix(words[i], "(up,") {
						for j := i - number; j < i; j++ {
							words[j] = strings.ToUpper(words[j])
						}
					} else if strings.HasPrefix(words[i], "(cap,") {
						for j := i - number; j < i; j++ {
							words[j] = strings.Title(words[j])
						}
					}
					words = append(words[:i], words[i+1:]...)
					i -= number
				}
			}
		}
	}
	return strings.Join(words, " ")
}
