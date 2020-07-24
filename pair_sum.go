package algo

import (
	"sort"
)

/*
	Given a list of numbers and a number k,
	return whether any two numbers from the list add up to k.
*/

// Sorting the array gives us a helfull property:
// Given an element in the array, the element to its right has a greater
// (or equal) value, and the element to its right has a lower (or equal value)
// (assing sort in ascending order)
// Using this property, we implement the "two-pointers technique"
//
// We place a pointer i on the the first array element, and a pointer j in the last
// We examine the value of the sum array[i] + array[j]
// If the sum is equal with k, return true
// if the sum is less than k, we need to "go closer to k", therefore move i one position to the right.
// Likewise, if the sum is less than k, move j one position to the left.
// Stop and return left when the pointers meet or cross.
func pairSumSort(array []int, k int) bool {
	if array == nil || len(array) == 0 {
		return false
	}

	sort.Ints(array)

	i := 0
	j := len(array) - 1

	for i < j {
		if array[i]+array[j] == k {
			return true
		} else if array[i]+array[j] < k {
			i++
		} else {
			j--
		}
	}

	return false
}

// A second solution is to use some additional memory in order to solve the problem
// doing a single pass on the array.
// We use a hash table to store, for each array element e, the information
// "does the array contain an element with value k-e?"
func pairSumHash(array []int, k int) bool {
	if array == nil || len(array) == 0 {
		return false
	}

	hMap := make(map[int]bool)
	for _, elem := range array {
		if _, ok := hMap[k-elem]; ok {
			return true
		}
		hMap[elem] = true
	}

	return false
}
