package hamming

import "errors"

func Distance(a, b string) (int, error) {
    if len(a) !=  len(b) {
        return 0, errors.New("The length of the given DSA strands are different") 
    }
    
	hamDistance := 0
    for i := range len(a) {
        if a[i] != b[i] {
            hamDistance += 1
        }
    }

    return hamDistance, nil
}
