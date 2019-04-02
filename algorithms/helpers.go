package algorithms

func DistanceBetweenXY(x, y float64) float64 {
	return x - y
}

func DistanceMultipliedBetweenXY(x, y float64) float64 {
	return x * y
}

func Similarity(number float64) float64 {
	return 1 / (1 + number)
}
