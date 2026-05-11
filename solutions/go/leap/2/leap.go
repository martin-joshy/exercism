// Package leap provides function IsLeapYear
package leap

// IsLeapYear takes in any given year and 
// returns boolean indicating whether it is 
// a leap year or not.
func IsLeapYear(year int) bool {
    if year % 400 == 0 {
        return true
    }
    if year % 4 == 0 && year % 100 != 0 {
        return true
    } 
    return false
}
