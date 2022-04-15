package homework4

func zeroThroughTen() []int {
	result := make([]int, 0, 10)
	// TODO: this loop currently never runs. Replace the loop header with values so that i
	// loops through [0, 10) in increasing order
	for i := 0; false; {
		result = append(result, i)
	}
	return result
}

func zeroThroughTenDecreasing() []int {
	result := make([]int, 0, 10)
	// TODO: this loop currently never runs. Replace the loop header with values so that i
	// loops through (10, 0] in decreasing order
	for i := 0; false; {
		result = append(result, i)
	}
	return result
}

func zeroThroughN(n int) []int {
	result := make([]int, 0, n)
	// TODO: this loop currently never runs. Replace the loop header with values so that i
	// loops through [0, n) in increasing order
	for i := 0; false; {
		result = append(result, i)
	}
	return result
}

func evensInRange(a, b int) []int {
	result := make([]int, 0)
	// TODO: this loop currently never runs. Replace the loop header with values so that i
	// loops through [a, b] in increasing order
	for i := 0; false; {
		// TODO: this if statement currently always executes. Replace the condition with a check
		// that i is even
		if true {
			result = append(result, i)
		}
	}
	return result
}

func sumThroughTen() int {
	sum := 0
	// TODO: this loop currently never runs. Replace the loop header with values so that i
	// loops through [1, 10] in increasing order
	for i := 0; false; {
		sum += i
	}
	return sum
}

func sumThroughN(n int) int {
	// TODO: implement this function on your own this time. Combine what you did in sumThroughTen
	// and zeroThroughN
	return 0
}

func productThroughTen() int {
	product := 1
	// TODO: this loop currently never runs. Replace the loop header with values so that i
	// loops through [1, 10] in increasing order
	for i := 0; false; {
		product *= i
	}
	return product
}

func productThroughN(n int) int {
	// TODO: implement this function on your own this time. Combine what you did in productThroughTen
	// and zeroThroughN
	return 0
}

func iterationsToN(n int) int {
	// TODO: implement this function
	return 0
}
