package geometry

import "math"

func Area(length, width float64) float64 {
	return length * width
}

// Diagonal function can be expoted
func Diagonal(length, width float64) float64 {
	return math.Sqrt((length * length) + (width * width))
}
