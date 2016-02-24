// Package disfun implements various distance functions.
//
package disfun

import (
	"github.com/gonum/matrix/mat64"
	"math"
)

const (
	EARTH_RADIUS = float64(6371)
	Substitution = float64(1)
	Insertion    = float64(1)
	Deletion     = float64(1)
)

// FloatsToMatrix creates a new mat64.Dense matrix from a slice of float64.
func FloatsToMatrix(floats []float64) *mat64.Dense {
	return mat64.NewDense(1, len(floats), floats)
}

// ZeroDense creates a new mat64.Dense matrix of zero values with dimenions of r and c.
func ZeroDense(r, c int) *mat64.Dense {
	var matrixPoints []float64
	for i := 0; i < (r * c); i++ {
		matrixPoints = append(matrixPoints, 0.0)
	}
	return mat64.NewDense(r, c, matrixPoints)
}

// MinInt32 finds the minimum int32 value of a range of int.
// MinInt32  = -1 << 31
func MinInt32(a ...int) int {
	min := math.MaxInt32
	for _, i := range a {
		if i < min {
			min = i
		}
	}
	return min
}

// MaxInt32 finds the maximum int32 value of a range of int.
// MaxInt32  = 1<<31 - 1
func MaxInt32(a ...int) int {
	max := math.MinInt32
	for _, i := range a {
		if i > max {
			max = i
		}
	}
	return max
}

// MinFloat64 finds the minimum float64 value of a range of float64.
// MaxFloat64 = 1.797693134862315708145274237317043567981e+308 // 2**1023 * (2**53 - 1) / 2**52
func MinFloat64(a ...float64) float64 {
	min := math.MaxFloat64
	for _, i := range a {
		if i < min {
			min = i
		}
	}
	return min
}

// MaxFloat64 finds the maximum int64 value of a range of float64.
// SmallestNonzeroFloat64 = 4.940656458412465441765687928682213723651e-324 // 1 / 2**(1023 - 1 + 52)
func MaxFloat64(a ...float64) float64 {
	max := math.SmallestNonzeroFloat64
	for _, i := range a {
		if i > max {
			max = i
		}
	}
	return max
}
