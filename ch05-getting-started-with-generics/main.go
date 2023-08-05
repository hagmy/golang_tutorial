package main

import "fmt"

func main() {
	ints := map[string]int64{
		"first":  34,
		"second": 50,
	}

	floats := map[string]float64{
		"first":  33.33,
		"second": 56.798,
	}

	fmt.Printf(
		"Non-Generic Sums: %v and %v\n",
		SumIntsOrFloats(ints),
		SumIntsOrFloats(floats),
	)
}

func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

// func SumInts(m map[string]int64) int64 {
// 	var s int64
// 	for _, v := range m {
// 		s += v
// 	}
// 	return s
// }

// func SumFloats(m map[string]float64) float64 {
// 	var s float64
// 	for _, v := range m {
// 		s += v
// 	}
// 	return s
// }
