package algorithms

func distanceBetweenXY(x, y float64)  float64 {
	return x - y
}

func distanceMultipliedBetweenXY(x, y float64)  float64 {
	return x * y
}

func similarity(number float64) float64{
	return 1 / (1 + number)
}
