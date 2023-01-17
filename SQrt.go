package main

import (
	"fmt"
	"math"
)

func SQrt(x float64) string {
	if x < 0 {
		return SQrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}
