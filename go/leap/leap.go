// Package leap provides a function to check if a year is a leap year
package leap

// IsLeapYear checks if a given year is a leap year
func IsLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}
