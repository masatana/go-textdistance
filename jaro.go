package textdistance

import (
	"math"
	"strings"
)

// JaroDistance calculates jaro distance between s1 and s2.
// This implementation is influenced by an implementation of [lucene](http://lucene.apache.org/)
// Note that this calculation's result is normalized ( the result will be bewtwen 0 and 1)
// and if t1 and t2 are exactly the same, the result is 1.0.
// This function returns distance and prefix (for jaro-winkler distance)
func JaroDistance(s1, s2 string) (float64, int) {
	if s1 == s2 {
		return 1.0, 0.0
	}
	var longer string
	var shorter string
	if len(s1) > len(s2) {
		longer = s1
		shorter = s2
	} else {
		longer = s2
		shorter = s1
	}
	scope := int(math.Floor(float64(len(longer)/2))) - 1
	// m is the number of matching characters.
	m := 0
	matchFlags := make([]bool, len(longer))
	matchIndexes := make([]int, len(longer))
	for i := range matchIndexes {
		matchIndexes[i] = -1
	}

	for i := 0; i < len(shorter); i++ {
		k := Min(i+scope+1, len(longer))
		for j := Max(i-scope, 0); j < k; j++ {
			if matchFlags[j] || shorter[i] != longer[j] {
				continue
			}
			matchIndexes[i] = j
			matchFlags[j] = true
			m++
			break
		}
	}
	ms1 := make([]uint8, m)
	ms2 := make([]uint8, m)
	si := 0
	for i := 0; i < len(shorter); i++ {
		if matchIndexes[i] != -1 {
			ms1[si] = shorter[i]
			si++
		}
	}
	si = 0
	for i := 0; i < len(longer); i++ {
		if matchFlags[i] {
			ms2[si] = longer[i]
			si++
		}
	}

	t := 0
	for i, c := range ms1 {
		if c != ms2[i] {
			t++
		}
	}
	prefix := 0
	longerArray := strings.Split(longer, "")
	shorterArray := strings.Split(shorter, "")
	for i := 0; i < len(shorter); i++ {
		if longerArray[i] == shorterArray[i] {
			prefix++
		} else {
			break
		}
	}
	if m == 0 {
		return 0.0, 0.0
	}
	newt := float64(t) / 2.0
	newm := float64(m)
	return 1 / 3.0 * (newm/float64(len(s1)) + newm/float64(len(s2)) + (newm-newt)/newm), prefix
}
