package gorandom

import "testing"

func TestIsVowel(t *testing.T) {
	var vowels = "aeiouAEIOU"
	var consonants = "bcdfghjklmnpqrstvwxyzBCDFGHJKLMNPQRSTVWXYZ"

	for _, element := range vowels {
		if !IsVowel(element) {
			t.Error("expected not err, got err")
		}
	}

	for _, element := range consonants {
		if IsVowel(element) {
			t.Error("expected not err, got err")
		}
	}
}

func TestIsTurkishVowel(t *testing.T) {
	var tr_vowels = "aeıioöuüAEIİOÖUÜ"
	var tr_consonants = "bcdfgğhjklmnprsştvyzBCDFGĞHJKLMNPRSŞTVYZ"

	for _, element := range tr_vowels {
		if !IsTurkishVowel(element) {
			t.Error("expected not err, got err when checking vowels")
		}
	}

	for _, element := range tr_consonants {
		if IsTurkishVowel(element) {
			t.Error("expected not err, got err when checking consonants")
		}
	}
}

func TestCountOfVowel(t *testing.T) {
	var words = []string{"strt", "car", "naMe", "AbUSe", "AUDIO", "euLogiA"}

	for index, element := range words {
		if index != CountOfVowel(element) {
			t.Error("expected not err, got err when checking count of vowels")
		}
	}
}

func TestCountOfTurkishVowel(t *testing.T) {
	var words = []string{"bşl", "zam", "bARan", "doĞuKan", "SüLEymANe", "Gündüşkünlügü"}

	for index, element := range words {
		if index != CountOfTurkishVowel(element) {
			t.Error("expected not err, got err when checking count of vowels")
		}
	}
}
