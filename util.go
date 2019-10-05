package textdistance

// Min returns the minimum number of passed int slices.
func Min(is ...int) int {
	var min int
	for i, v := range is {
		if i == 0 || v < min {
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
