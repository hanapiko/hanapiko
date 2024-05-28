package functions

import "math"

func StandardDev(num []float64) float64{
	variance := Variance(num)
	return math.Sqrt(variance)
}