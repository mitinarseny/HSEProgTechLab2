package search

func Full(n int, f func(int) bool) int {
	for i := 0; i < n; i++ {
		if f(i) {
			return i
		}
	}
	return n
}
