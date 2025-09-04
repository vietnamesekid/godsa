package algorithms

import "slices"

func LinearSearch(arr []int, needle int) bool {
	return slices.Contains(arr, needle)
}
