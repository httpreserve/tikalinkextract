package main

//math.Min uses float64, so let's not cast
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
