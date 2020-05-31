package main

import (
	"sort"
)

// Loop with O(n*n)
func solution1(a []int, b int) []int {
	n := len(a)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if a[i]+a[j] == b {
				if a[i] <= a[j] {
					return []int{a[i], a[j]}
				} else {
					return []int{a[j], a[i]}
				}

			}
		}

	}
	return []int{}
}

func sortInput(a []int) []int {
	sort.Slice(a, func(i, j int) bool {
		return a[i] <= a[j]
	})

	return a
}

// O(n*logN) + O(n)
func solution2(a []int, b int) []int {
	n := len(a)

	a = sortInput(a)

	for i, j := 0, n-1; i <= j; {
		if a[i]+a[j] == b {
			return []int{a[i], a[j]}
		} else if a[i]+a[j] > b {
			j--
		} else {
			i++
		}
	}

	return []int{}
}

// Using a Map with O(n)
func solution3(a []int, b int) []int {
	n := len(a)
	checker := map[int]bool{}

	for i := 0; i < n; i++ {
		if checker[b-a[i]] {
			if a[i] <= b-a[i] {
				return []int{a[i], b - a[i]}
			} else {
				return []int{b - a[i], a[i]}
			}
		} else {
			checker[a[i]] = true
		}
	}

	return []int{}
}

func main() {
	a := []int{3, 5, -4, 8, 11, 1, -1, 6}
	b := 10

	result := solution3(a, b)
	for _, v := range result {
		println(v)
	}
}
