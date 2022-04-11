package homework3

func subtract(x, y int) int {
	return x - y
}

func multiply(x, y int) int {
	return x * y
}

func divide(x, y int) int {
	return x / y
}

func sum(x, y, z int) int {
	return x + y + z
}

func isEven(x int) bool {
	return x%2 == 0
}

func isOdd(x int) bool {
	return x%2 == 1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func and(x, y bool) bool {
	return x && y
}

func or(x, y bool) bool {
	return x || y
}

func not(x bool) bool {
	return !x
}
