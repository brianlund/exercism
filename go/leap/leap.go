// Package leap solves the leap side exercise on exercism.io
package leap

// IsLeapYear take a year as input and returns if it is a leap year or not
func IsLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}
