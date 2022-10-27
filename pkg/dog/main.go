// Allowing to understand more about dog
package dog

// Converting dog years to human years.
func Years(h int) int {
	return h * 7
}

// Converting dog months to human months.
func Month(h int) int {
	return Years(h) * 12
}
