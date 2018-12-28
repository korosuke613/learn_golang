package plactice

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	old := 0
	oldOld := 1
	return func() int {
		result := old + oldOld
		oldOld = old
		old = result
		return result
	}
}
