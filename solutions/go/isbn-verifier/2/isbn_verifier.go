package isbnverifier
import "strings"

func IsValidISBN(isbn string) bool {
    isbn = strings.ReplaceAll(isbn, "-", "")
    if len(isbn) != 10 {
        return false
    }
    
    sum := 0
    for pos, r := range isbn{
        var currVal int
              
        if pos == 9 && r == 'X' {
            currVal = 10
        }else if r >= '0' && r <= '9' {
            currVal = int(r - '0')
        }else{
            return false
        }
        
        sum += currVal * (10 - pos)
    }

    return sum % 11 == 0
}