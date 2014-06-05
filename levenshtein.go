package textdistance

import "strings"

func LevenshteinDitance(s1, s2 string) int {
	if s1 == s2 {
		return 0
	}
	s1Array := strings.Split(s1, "")
	s2Array := strings.Split(s2, "")
	lenS1Array := len(s1Array)
	lenS2Array := len(s2Array)
	if lenS1Array == 0 {
		return lenS2Array
	}
	if lenS2Array == 0 {
		return lenS1Array
	}
	m := make([][]int, lenS1Array+1)
	dis := 0
	for i := range m {
		m[i] = make([]int, lenS2Array+1)
	}
	for i := 0; i < lenS1Array+1; i++ {
		for j := 0; j < lenS2Array+1; j++ {
			if Min(i, j) == 0 {
				m[i][j] = Max(i, j)
			} else {
				dis = 0
				if s1Array[i-1] != s2Array[j-1] {
					dis = 1
				}
				m[i][j] = Min(m[i-1][j]+1, m[i][j-1]+1, m[i-1][j-1]+dis)
			}
		}
	}
	return m[lenS1Array][lenS2Array]
}
