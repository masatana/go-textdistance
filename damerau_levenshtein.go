package textdistance

func DamerauLevenshteinDistance(s1, s2 string) int {
	lenS1 := len(s1)
	lenS2 := len(s2)
	m := make([][]int, lenS1+1)
	cost := 0
	for i := range m {
		m[i] = make([]int, lenS2+1)
	}
	for i := 0; i < lenS1+1; i++ {
		for j := 0; j < lenS2+1; j++ {
			if Min(i, j) == 0 {
				m[i][j] = Max(i, j)
			} else {
				cost = 0
				if s1[i-1] != s2[j-1] {
					m[i][j] = Min(m[i-1][j]+1, m[i][j-1]+1, m[i-1][j-1]+cost)
				}
				if i > 1 && j > 1 && s1[i-1] == s2[j-2] && s1[i-2] == s1[j-1] {
					m[i][j] = Min(m[i][j], m[i-3][j-3]+cost)
				}
			}
		}
	}
	return m[lenS1][lenS2]
}
