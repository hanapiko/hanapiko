package functions

func Median(num []float64) float64{
	value := len(num)
	if num == nil {
		return 0.0
	}
	if value %2 == 0{
		return (num[value/2-1]+num[value/2])/2
	} else {
		return num[value/2]
	}
}