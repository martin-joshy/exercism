// Package leap provides function IsLeapYear
package leap

// IsLeapYear takes in any given year and 
// returns boolean indicating whether it is 
// a leap year or not.
func IsLeapYear(year int) bool {
    return year % 4 == 0 && (year % 100 !=0 || year % 400 ==0)
}
