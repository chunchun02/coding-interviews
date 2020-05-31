package main

// Recursion with O(2^n)
func solution1(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	return solution1(n-1) + solution1(n-2)
}

// Loop with O(n) in time complexity and O(1) extra space
func solution2(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	a, b, c := 0, 1, 1

	for i := 2; i <= n; i++ {
		a = b
		b = c
		c = a + b
	}

	return c
}

func main() {
	n := 10

	print(solution2(n))
}
