package anagram

import (
    "maps"
    "unicode"
    "strings"
)

func Detect(subject string, candidates []string) []string {
	subjCount := make(map[rune]int, len(subject))
    res := make([]string, 0, len(candidates))

    for _, r := range subject {
        subjCount[unicode.ToLower(r)]++
    }  
    
    for _, cndt := range candidates {
        if strings.ToLower(cndt) == strings.ToLower(subject) {
            continue
        }
        cndtCount := make(map[rune]int)
        
        for _, r := range cndt {
            cndtCount[unicode.ToLower(r)]++
        }
        if maps.Equal(cndtCount, subjCount){
            res = append(res, cndt)
        }
    }
    
    return res
}
