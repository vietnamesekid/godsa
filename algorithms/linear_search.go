package algorithms

func LinearSearch(arr []int, needle int) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] == needle {
			return true
		}
	}

	return false
}
