package gorandom

// Is letter vowel check
func IsVowel(char rune) bool {
	switch {
	case
		char == 'a',
		char == 'e',
		char == 'i',
		char == 'u',
		char == 'o':
		return true
	default:
		return false
	}
}

// Is letter Turkish vowel check
func IsTurkishVowel(char rune) bool {
	switch {
	case
		char == 'a',
		char == 'e',
		char == 'ı',
		char == 'i',
		char == 'u',
		char == 'ü',
		char == 'o',
		char == 'ö':
		return true
	default:
		return false
	}
}

// Count of vowel in a string
func CountOfVowel(string_ string) int {
	var count = 0
	for _, character := range string_ {
		if IsVowel(character) {
			count += 1
		}
	}

	return count
}

// Count of Turkish vowel in a string
func CountOfTurkishVowel(string_ string) int {
	var count = 0
	for _, character := range string_ {
		if IsTurkishVowel(character) {
			count += 1
		}
	}

	return count
}
