package samples

func keys(m map[string]bool) (result []string) {
	result = make([]string, 0, len(m))
	for key := range m {
		result = append(result, key)
	}
	return
}

func removeDups(a []string) []string {
	var m map[string]bool = make(map[string]bool, len(a))
	for _, val := range a {
		_, found := m[val]
		if !found {
			m[val] = false
		} else {
			m[val] = true
		}
	}
	return keys(m)
}

func Merge(a, b []string) []string {
	return removeDups(append(a, b...))
}
