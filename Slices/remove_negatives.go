package slices

// RemoveNegatives entfernt alle negativen Zahlen aus dem Slice,
// indem es das Slice direkt über den übergebenen Pointer verändert.
func RemoveNegatives(nums []int) []int {
var num []int

	for _,v := range nums {
	
		if v >= 0 {
			num = append(num, v)
		}
	}

	return num
}
