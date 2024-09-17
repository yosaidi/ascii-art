package main

func AreStringValid(runes []rune) (bool, rune) {
	for _, char := range runes {
		if (char < 32 || char > 126) && char != '\n' {
			return false, char
		}
	}
	return true, 0
}