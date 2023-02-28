package util

import "fmt"

func SliceContains[T comparable](xs []T, y T) bool {
	for _, x := range xs {
		if x == y {
			return true
		}
	}

	return false
}

func ConvertToStrings[T fmt.Stringer](stringers []T) []string {
	values := make([]string, 0, len(stringers))
	for _, v := range stringers {
		values = append(values, v.String())
	}

	return values
}
