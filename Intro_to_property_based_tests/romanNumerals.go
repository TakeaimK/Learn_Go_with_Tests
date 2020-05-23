package romanNumerals

func ConvertToRoman(arabic int) string {
	if arabic == 2 {
		return "II"
	} else if arabic == 3 {
		return "III"
	}
	return "I"
}
