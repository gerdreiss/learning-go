package main

// Sum sums the values of map m. It supports both int64 and float64
// as types for map values.
func Sum[K comparable, V int64 | uint64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}
