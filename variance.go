package functions

//"mathskills/functions"

func Variance(num []float64) float64{
	sum := 0.0
	//average := functions.Average(num)
	average := Average(num)
	for _, val := range num{
	sum += (val-average) * (val - average)
}
return sum/float64(len(num))
}