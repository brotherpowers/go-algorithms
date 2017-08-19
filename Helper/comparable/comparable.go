package comparable

// Comparable interface
type Comparable interface {
	// -1 Lesser | 1 Greater | 0 Equals
	Compare(Comparable) int
}
