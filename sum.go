package gofunc

// Sum : Adds together all the elements of an int list.
func Sum(list []int) int {
	result := 0
	for _, value := range list {
		result += value
	}
	return result
}
