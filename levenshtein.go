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

func LevenshteinDitance(s1, s2 string) int {
	lenS1 := len(s1)
	lenS2 := len(s2)
	m := make([][]int, lenS1+1)
	dis := 0
	for i := range m {
		m[i] = make([]int, lenS2+1)
	}
	for i := 0; i < lenS1+1; i++ {
		for j := 0; j < lenS2+1; j++ {
			if Min(i, j) == 0 {
				m[i][j] = Max(i, j)
			} else {
				dis = 0
				if s1[i-1] != s2[j-1] {
					dis = 1
				}
				m[i][j] = Min(m[i-1][j]+1, m[i][j-1]+1, m[i-1][j-1]+dis)
			}
		}
	}
	return m[lenS1][lenS2]
}
