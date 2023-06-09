package ch06

func mean(vals []float64) float64 {
	var sum float64
	for _, val := range vals {
		sum += val
	}
	return sum / float64(len(vals))
}

func variance(vals []float64) float64 {
	m := mean(vals)
	var sum float64
	for _, val := range vals {
		sum += (val - m) * (val - m)
	}
	return sum / float64(len(vals))
}

// mean square error
func mse(vals []float64, preds []float64) float64 {
	var sum float64
	for i, val := range vals {
		sum += (val - preds[i]) * (val - preds[i])
	}
	return sum / float64(len(vals))
}
