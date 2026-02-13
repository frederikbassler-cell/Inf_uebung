package slices

// RemoveNegatives entfernt alle negativen Zahlen aus dem Slice,
// indem es das Slice direkt über den übergebenen Pointer verändert.
func RemoveNegativesRekursiv(nums []int) []int {
	var temp []int
	remove(nums, 0, &temp)				
	return temp
}

func remove(nums []int, count int, temp *[]int) int {
	if count >= len(nums) {
		return 0
	}

	if nums[count] >= 0 {
		*temp = append(*temp, nums[count])
	}

	remove(nums, count+1, temp)	
	return 0
}
