package roman_numerals

import "strings"

type RomanNumerals struct {
	Value  int
	Symbol string
}

var allRomanNumerals = []RomanNumerals{
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ConvertToRoman(arabic int) string {
	var result strings.Builder
	for _, romanNumeral := range allRomanNumerals {
		if arabic >= romanNumeral.Value {
			arabic -= romanNumeral.Value
			result.WriteString(romanNumeral.Symbol)
		}
	}
	return result.String()
}
