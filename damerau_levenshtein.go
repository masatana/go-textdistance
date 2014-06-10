package textdistance

import "strings"

func DamerauLevenshteinDistance(s1, s2 string) int {
	if s1 == s2 {
		return 0
	}
	s1Array := strings.Split(s1, "")
	s2Array := strings.Split(s2, "")
	lenS1Array := len(s1Array)
	lenS2Array := len(s2Array)
	m := make([][]int, lenS1Array+1)
	var cost int
	for i := range m {
		m[i] = make([]int, lenS2Array+1)
	}
	for i := 0; i < lenS1Array+1; i++ {
		for j := 0; j < lenS2Array+1; j++ {
			if i == 0 {
				m[i][j] = j
			} else if j == 0 {
				m[i][j] = i
			} else {
				cost = 0
				if s1Array[i-1] != s2Array[j-1] {
					cost = 1
				}
				m[i][j] = Min(m[i-1][j]+1, m[i][j-1]+1, m[i-1][j-1]+cost)
				if i > 1 && j > 1 && s1Array[i-1] == s2Array[j-2] && s1Array[i-2] == s2Array[j-1] {
					m[i][j] = Min(m[i][j], m[i-2][j-2]+cost)
				}
			}
		}
	}
	return m[lenS1Array][lenS2Array]
}
