package data_info

import "math"

var MaxCount = 400

var SinFloat64 = (func() []float64 {
	data := make([]float64, MaxCount)
	for i := range data {
		data[i] = 1 + math.Sin(float64(i)/5)
	}
	return data
})()
