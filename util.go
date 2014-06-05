package textdistance

// Min returns the minimum number of passed int slices.
func Min(is ...int) int {
	min := int(1e10)
	for _, v := range is {
		if min > v {
			min = v
		}
	}
	return min
}

// Max returns the maximum number of passed int slices.
func Max(is ...int) int {
	var max int
	for _, v := range is {
		if max < v {
			max = v
		}
	}
	return max
}
