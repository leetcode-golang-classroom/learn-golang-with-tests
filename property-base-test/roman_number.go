package roman_numerals

import "strings"

func ConvertToRoman(arabic int) string {
	var result strings.Builder
	for arabic > 0 {
		switch {
		case arabic > 4:
			arabic -= 5
			result.WriteString("V")
		case arabic > 3:
			arabic -= 4
			result.WriteString("IV")
		default:
			arabic--
			result.WriteString("I")
		}
	}
	return result.String()
}
