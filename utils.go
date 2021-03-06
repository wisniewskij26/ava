package ava

func average(x interface{}) float64 {
	switch t := x.(type) {
	case []float64:
		return averageOfFloats(t)
	case [][]float64:
		return averageofNestedFloats(t)
	}
	return 0
}

func averageOfFloats(x []float64) float64 {
	return sum(x) / float64(len(x))
}

func averageofNestedFloats(x [][]float64) float64 {
	var total float64
	var n int
	for _, xi := range x {
		total += sum(xi)
		n += len(xi)
	}
	return total / float64(n)
}

func sum(x []float64) float64 {
	var sum float64
	for _, xi := range x {
		sum += xi
	}
	return sum
}

func nestedLen(groups [][]float64) int {
	var n int
	for _, j := range groups {
		n += len(j)
	}
	return n
}

func variance(sample []float64) float64 {
	var sumOfSquaredDiff float64
	xbar := average(sample)

	for _, xi := range sample {
		sumOfSquaredDiff += (xi - xbar) * (xi - xbar)
	}
	return sumOfSquaredDiff / float64(len(sample)-1)
}

func pooledVariance(groups [][]float64) float64 {
	var numer float64
	var denom float64
	for _, sample := range groups {
		numer += float64(len(sample)-1) * variance(sample)
		denom += float64(len(sample) - 1)
	}
	return numer / denom
}
