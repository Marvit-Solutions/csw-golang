package helper

func CleanNumericString(s string) string {
	var cleaned string
	for _, char := range s {
		if char >= '0' && char <= '9' {
			cleaned += string(char)
		}
	}
	return cleaned
}
