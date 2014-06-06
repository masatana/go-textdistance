package textdistance

func JaroWinklerDistance(s1, s2 string) float64 {
	threshold := 0.7
	jaroDistance, prefix := JaroDistance(s1, s2)
	if jaroDistance < threshold {
		return jaroDistance
	}
	return jaroDistance + (float64(Min(prefix, 4)) * 0.1 * (1 - jaroDistance))
}
