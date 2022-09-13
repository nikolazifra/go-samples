package samples

func keys[K comparable, V any](m map[K]V) (result []K) {
	result = make([]K, 0, len(m))
	for key := range m {
		result = append(result, key)
	}
	return
}

func removeDups[K comparable](a []K) []K {
	var m map[K]bool = make(map[K]bool, len(a))
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

func Merge[K comparable](a, b []K) []K {
	return removeDups(append(a, b...))
}
