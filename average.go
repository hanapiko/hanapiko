package functions

func Average(num []float64) float64{
	sum := 0.0
	for _, v := range num {
        sum += v
    }
	return sum / float64(len(num))
}