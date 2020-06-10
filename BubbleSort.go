package process

//	@BubbleSort
// 	@description  	[]int Follow small to big
// 	@Author 		CyBen
// 	@param   		[]int{} 		[]int
//	@return   		[]int{} 		[]int
func BubbleSortAscending(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		for j := 1; j < len(nums)-i; j++ {
			if nums[j] < nums[j-1] {
				//交换
				nums[j], nums[j-1] = nums[j-1], nums[j]
			}
		}
	}
	return nums
}

//	@BubbleSort
// 	@description  	[]int Follow big to small
// 	@Author 		CyBen
// 	@param   		[]int{} 		[]int
//	@return   		[]int{} 		[]int
func BubbleSortdescending(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		for j := 1; j < len(nums)-i; j++ {
			if nums[j] > nums[j-1] {
				//交换
				nums[j], nums[j-1] = nums[j-1], nums[j]
			}
		}
	}
	return nums
}
