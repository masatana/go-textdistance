package textdistance

func Min(is ...int) int {
	min := 100
	for _, v := range is {
		if min > v {
			min = v
		}
	}
	return min
}

func Max(is ...int) int {
	var max int
	for _, v := range is {
		if max < v {
			max = v
		}
	}
	return max
}
