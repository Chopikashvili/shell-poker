package general

import "slices"

// Filters elements from a slice according to predicate.
func Filter[T any](slice []T, predicate func(T) bool) []T {
	subslice := make([]T, 0)
	for _, elem := range slice {
		if predicate(elem) {
			subslice = append(subslice, elem)
		}
	}
	return subslice
}

// Counts certain numbers in an int slice.
func Count(slice []int, comparable int) int {
	var counter int
	for _, elem := range slice {
		if elem == comparable {
			counter++
		}
	}
	return counter
}

// Checks if the first slice contains a subslice that is equal to the second slice.
func ContainsSubslice[T comparable](slice []T, subslice []T) bool {
	for i := 0; i < len(slice)-len(subslice)+1; i++ {
		if slices.Equal(slice[i:i+len(subslice)], subslice) {
			return true
		}
	}
	return false
}

// Sorts int arrays in reverse, returns a copy of the original and the result.
func SortDesc(s []int) ([]int, []int) {
	orig := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		orig[i] = s[i]
	}
	slices.Sort(s)
	slices.Reverse(s)
	return orig, s
}

func CopySlice[T any](slice []T) []T {
	copy := []T{}
	for _, elem := range slice {
		copy = append(copy, elem)
	}
	return copy
}
