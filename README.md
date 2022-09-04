## go-random | Random String Generator <img src="https://img.shields.io/github/license/zekiahmetbayar/go-random?style=for-the-badge" width="100" />
<p float="left">
  <img src="https://img.shields.io/github/downloads/zekiahmetbayar/go-random/total?style=for-the-badge" width="150" />
  <img src="https://img.shields.io/github/commit-activity/m/zekiahmetbayar/go-random?style=for-the-badge" width="170" /> 
</p>


The gorandom package provides random string, hex string, vowel check and count of vowel functions.

### Installation

```
go get -u "github.com/zekiahmetbayar/go-random"
```

### Usage

```go

package main

import (
	"fmt"
	"log"

	gorandom "github.com/zekiahmetbayar/go-random"
)

func main() {
	// Returns string containing numbers, letters and special characters of given length
	r_str, err := gorandom.String(true, true, true, 16)
	if err != nil {
		log.Fatalf("error when creating random string: %s", err.Error())
	}
	fmt.Printf("Random String: %s\n", r_str)

	// Returns hex string of given length
	h_str, err := gorandom.HexString(16)
	if err != nil {
		log.Fatalf("error when creating hex string: %s", err.Error())
	}
	fmt.Printf("Hex String: %s\n", h_str)

	// Returns count of vowels in string
	count := gorandom.CountOfVowel("test")
	fmt.Printf("Count: %d\n", count)

	// Returns count of Turkish vowels in string
	t_count := gorandom.CountOfTurkishVowel("ali")
	fmt.Printf("Count: %d\n", t_count)

	// Returns if given character is vowel
	for _, element := range "isvowel" {
		is_vowel := gorandom.IsVowel(element)
		fmt.Printf("Is vowel: %v\n", is_vowel)
	}

	// Returns if given character is Turkish vowel
	for _, element := range "türkçe ünlü karakter mi" {
		is_Tvowel := gorandom.IsTurkishVowel(element)
		fmt.Printf("Is Turkish vowel: %v\n", is_Tvowel)
	}
}

```