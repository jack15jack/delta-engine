package indicators

func SMA(values []float64, period int) float64 {

	if len(values) < period {
		return 0
	}

	sum := 0.0

	for i := len(values) - period; i < len(values); i++ {
		sum += values[i]
	}

	return sum / float64(period)
}
