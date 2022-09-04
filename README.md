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

import "github.com/zekiahmetbayar/go-random"

func main() {
    // Returns string containing numbers, letters and special characters of given length
    r_str, err := String(true, true, true, 16)
    if err != nil {
        log.Fatalf("error when creating random string: %s", err.Error())
    }
    fmt.Printf("Random String: %s\n", r_str)

    // Returns hex string of given length
    h_str, err := HexString(16)
    if err != nil {
        log.Fatalf("error when creating hex string: %s", err.Error())
    }
    fmt.Printf("Hex String: %s\n", h_str)

    // Returns count of vowels in string
    count, err := CountOfVowel("test")
    if err != nil {
        log.Fatalf("error when calculating count of vowels in string: %s", err.Error())
    }
    fmt.Printf("Count: %d\n", count)

    // Returns count of Turkish vowels in string
    t_count, err := CountOfTurkishVowel("ali")
    if err != nil {
        log.Fatalf("error when calculating count of Turkish vowels in string: %s", err.Error())
    }
    fmt.Printf("Count: %d\n", t_count)

    // Returns if given character is vowel
    is_vowel, err := IsVowel("X")
    if err != nil {
        log.Fatalf("error when checking if given character is vowel: %s", err.Error())
    }
    fmt.Printf("Is vowel: %v\n", is_vowel)

    // Returns if given character is Turkish vowel
    is_Tvowel, err := IsTurkishVowel("Ü")
    if err != nil {
        log.Fatalf("error when checking if given character is Turkish vowel: %s", err.Error())
    }
    fmt.Printf("Is Turkish vowel: %v\n", is_Tvowel)
}

```