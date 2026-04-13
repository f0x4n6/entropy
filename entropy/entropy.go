// Package entropy provides methods to calculate entropy of block.
//
// Source: https://gist.github.com/n2p5/4eda328b080c9f09eff928ad47228ab1
package entropy

import "math"

// Calculate returns the entropy of the given block.
func Calculate(block []byte) float64 {
	var n = float64(len(block))
	var a = make([]float64, 256)
	var e float64

	for _, b := range block {
		a[b]++
	}

	for i := range 256 {
		if a[i] != 0 {
			v := a[i] / n
			e -= v * math.Log2(v)
		}
	}

	return e
}
