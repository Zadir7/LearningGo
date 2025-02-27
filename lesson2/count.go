package lesson2

func Count(s string, r rune) int {
	count := 0
	for _, char := range s {
		if char == r {
			count++
		}
	}
	return count
}
