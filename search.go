package search

//Binary ...
func Binary(arrayParams []int, start int, totalRecord int, searchVariable int) int {
	totalRecord = totalRecord - 1
	for start <= totalRecord {
		mid := start + (totalRecord-start)/2
		if arrayParams[mid] == searchVariable {
			return mid
		}
		if arrayParams[mid] < searchVariable {
			start = mid + 1
		} else {
			totalRecord = mid - 1
		}
	}
	return -1
}

//Bubble
func Bubble() {

}
