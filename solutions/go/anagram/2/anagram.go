package anagram

import (
    "maps"
    "strings"
)

func Detect(subject string, candidates []string) []string {
	subjCount := make(map[rune]int, len(subject))
    res := make([]string, 0, len(candidates))

    subjLower := strings.ToLower(subject)
    for _, r := range subjLower {
        subjCount[r]++
    }  
    
    for _, cndt := range candidates {
        cndtLower := strings.ToLower(cndt)
        if cndtLower == subjLower {
            continue
        }
        
        cndtCount := make(map[rune]int)
        for _, r := range cndtLower {
            cndtCount[r]++
        }
        if maps.Equal(cndtCount, subjCount){
            res = append(res, cndt)
        }
    }
    
    return res
}
