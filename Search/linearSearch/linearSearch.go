package linearSearch

// Search algo
func Search(src []interface{}, searchable interface{}) int {
	if len(src) > 0 {
		for i, element := range src {
			if searchable == element {
				return i
			}
		}
	}
	return -1
}
