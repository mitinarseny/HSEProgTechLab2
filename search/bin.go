package search

func Bin(n int, f func(int) bool) int {
	i, j := 0, n
	for i < j {
		h := int(uint(i + j) >> 1)
		if f(h) {
			j = h
		} else {
			i = h + 1
		}
	}
	return i
}