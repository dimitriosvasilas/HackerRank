package algo

/*
	Given an array of integers, return a new array such that
	each element at index i of the new array is the product of all the numbers
	in the original array except the one at i.
*/

// A simple approach.
// Do a first pass on the array, calculating the product of all elements (p)
// Then, do a second pass, and for each element e, calculate p/e
func productOfRestDiv(array []int) []int {
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

// Can we solve it without using division?
// We need a way a constructive solution.
// The insight is that each element splits the array into the elements before it (prefix),
// and the ones after it (suffix).
// So constructing the required products can be broken down to:
// - calculating the product of each element's array prefix,
// - calculating the product of each element's array suffix,
// - multiplying the two.
// To achieve that, we can use an array that, for each element, stores the product of its array prefix
// and one that stores its array suffix,
// and then multiplying those arrays per-element.
// Here, we merge both arrays in one.
func productOfRest(array []int) []int {
	if array == nil {
		return nil
	}

	if len(array) == 0 {
		return []int{}
	}

	resp := make([]int, len(array))
	p := 1
	for i := 0; i < len(array); i++ {
		resp[i] = p
		p *= array[i]
	}

	p = 1
	for i := len(array) - 1; i >= 0; i-- {
		resp[i] *= p
		p *= array[i]
	}

	return resp
}
