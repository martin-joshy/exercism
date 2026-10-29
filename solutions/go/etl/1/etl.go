package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	out := make(map[string]int)

    for pts, chars := range in {
        for _, char := range chars {
            out[strings.ToLower(char)] = pts
        } 
    }

    return out
}
