package main

func AreStringValid(runes []rune) (bool, rune) {
	for _, char := range runes {
		if (char < 32 || char > 126) && char != '\n' {
			return false, char
		}
	}
	return true, 0
}


func SplitWith(str string) []string {
	myStr := str
	result := []string{}
	helper := ""
	for i := 1; i <len(myStr); i++{
		if myStr[i] == 'n' && myStr[i-1] == '\\' {
			result = append(result, helper)
		} else {
			helper += string(myStr[i])
		}
	}
	return result
}