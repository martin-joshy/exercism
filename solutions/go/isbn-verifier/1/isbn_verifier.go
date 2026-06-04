package isbnverifier

func IsValidISBN(isbn string) bool {
    sum := 0
    pos := 0

    for _, r := range isbn{
        if r == '-'{
            continue
        }
        if pos == 9 && r == 'X' {
            sum += 10
            pos++
            continue
        }
        if r < '0' || r > '9' {
            return false
        }
        sum += int(r - '0') * (10 - pos)
        pos++
    }

    return pos == 10 && sum % 11 == 0
}


