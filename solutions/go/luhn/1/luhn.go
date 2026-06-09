package luhn

func Valid(id string) bool {
    checkSum := 0
    digitCount := 0
	for i :=len(id) - 1; i>= 0; i--{
        if id[i] == ' '{
            continue
        }
        if id[i] < '0' || id[i] > '9'{
            return false
        }
        digit := int(id[i] - '0') 
        if digitCount % 2 == 1 {
            digit *= 2
            if digit > 9 {
                digit -= 9
            }
        }
        
        checkSum += digit
        digitCount++
    }
    return digitCount > 1 && checkSum % 10 == 0
}
