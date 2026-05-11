package hamming

import "errors"

func Distance(a, b string) (int, error) {
    if len(a) !=  len(b) {
        return 0, errors.New("The length of the given DNA strands are different") 
    }
    
	hamDistance := 0
    for i, rn := range a {
        if byte(rn) != b[i] {
            hamDistance++
        }
    }

    return hamDistance, nil
}
