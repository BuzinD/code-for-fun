package main

import "fmt"

func main() {
	source := []int64{2, 4, 6, 8, 10, 12}
	var target int64 = 8

	fmt.Println(source, target, binSearch(source, target))
}

func binSearch(source []int64, target int64) int {
	l := 0
	r := len(source) - 1

	for l <= r {
		mid := (l + r) / 2

		if source[mid] == target {
			return mid
		}

		if source[mid] > target {
			r = mid - 1
		} else if source[mid] < target {
			l = mid + 1
		}
	}

	return -1
}
