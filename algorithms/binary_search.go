package algorithms

func BinarySearch(arr []int, needle int) bool {
	if len(arr) == 0 {
		return false
	}
	left, right := 0, len(arr)-1
	for left <= right {
		mid := left + (right-left)>>1 // use bit shift for division by 2
		switch {
		case arr[mid] == needle:
			return true
		case arr[mid] < needle:
			left = mid + 1
		default:
			right = mid - 1
		}
	}
	return false
}
