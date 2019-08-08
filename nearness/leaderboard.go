package nearness

func BestRawScores() map[int]int {
	// TODO: fetch this from the website http://azspcs.com/Contest/Nearness/BestRawScores.
	return map[int]int{
		6:  5526,
		7:  17779,
		8:  57152,
		9:  144459,
		10: 362950,
		11: 740798,
		12: 1585264,
		13: 2888120,
		14: 5457848,
		15: 9164700,
		16: 15891088,
		17: 25152826,
		18: 40901354,
		19: 61784724,
		20: 95128264,
		21: 138135156,
		22: 203913844,
		23: 286970392,
		24: 409184334,
		25: 560456128,
		26: 776393380,
		27: 1039513584,
		28: 1404896640,
		29: 1843448171,
		30: 2439470876,
	}
}

func TotalNormalizedScore(solutions map[int]*Board) float64 {
	var sum float64
	for _, b := range solutions {
		sum += NormalizedScore(b)
	}
	return sum
}

func NormalizedScore(b *Board) float64 {
	rawScores := BestRawScores()
	return float64(rawScores[b.Size]) / float64(b.NormNearness())
}
