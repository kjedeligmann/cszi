package main

func countLetters(text string) []int {
	freqs := make([]int, ukrAlfLen)
	for _, rune := range text {
		if char, in := ukrLetterToIdx[rune]; in {
			freqs[char.Idx]++
		}
	}
	return freqs
}

var binLogProbUkr = []float64{-3.6, -6.0, -4.2, -6.0, -13.3, -4.8, -4.4, -8.0, -7.1, -5.6, -4.1, -6.3, -4.0, -6.9, -4.6, -4.7, -5.0, -3.8, -3.4, -5.1, -4.2, -4.5, -4.4, -4.9, -8.2, -6.4, -6.6, -6.4, -7.1, -8.3, -5.8, -7.2, -5.3}

func logLikelihood(freqs []int, key int) (sum float64) {
	for i := 0; i < ukrAlfLen; i++ {
		sum += float64(freqs[mod(i+key, ukrAlfLen)]) * binLogProbUkr[i]
	}
	return sum
}

func llForKeys(freqs []int) ([]float64, int) {
	var maxLL int
	logLikelihoods := make([]float64, ukrAlfLen)

	for key := 0; key < ukrAlfLen; key++ {
		logLikelihoods[key] = logLikelihood(freqs, key)

		if logLikelihoods[key] > logLikelihoods[maxLL] {
			maxLL = key
		}
	}
	return logLikelihoods, maxLL
}
