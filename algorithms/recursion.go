package algorithms

func Sum(n int) int {
	if n == 1 {
		return 1
	}

	return n + Sum(n-1)
}
