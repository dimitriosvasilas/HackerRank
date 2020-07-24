package algo

// Given an array of integers, return a new array such that
// each element at index i of the new array is the product of all the numbers
// in the original array except the one at i.

// -------

// A simple approach.
// Do a first pass on the array, calculating the product of all elements (p)
// Then, do a second pass, and for each element e, calculate p/e
func productOfRest(array []int) []int {
	if array == nil {
		return nil
	}

	if len(array) == 0 {
		return []int{}
	}

	resp := make([]int, len(array))
	p := 1
	for _, elem := range array {
		p *= elem
	}

	for i, elem := range array {
		resp[i] = p / elem
	}

	return resp
}
