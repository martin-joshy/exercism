package pangram

import (
    "strings"
    "unicode"
    )

func IsPangram(input string) bool {
	freq := [26]int{}

    for _, r := range strings.ToLower(input) {
        if unicode.IsLetter(r) {
            freq[r - 'a']++
        }
    }

    for _, count := range freq {
        if count == 0 {
            return false
        }
    }

    return true
}
