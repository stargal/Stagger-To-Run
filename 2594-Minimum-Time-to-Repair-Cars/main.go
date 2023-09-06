package main

import "math"

// https://leetcode.cn/problems/minimum-time-to-repair-cars/

// 需要转换思维，不要从修这么多车需要多少时间入手（遍历，复杂度爆炸），要从这么多时间能修多少车入手（对时间进行二分查找）。

func main() {
	ranks := []int{3, 3, 1, 2, 1, 1, 3, 2, 1}
	cars := 58
	println(repairCars(ranks, cars))
}

func repairCars(ranks []int, cars int) int64 {
	left, right := 1, ranks[0]*cars*cars

	var check = func(time int64) bool {
		var total int64
		for _, rank := range ranks {
			total += int64(math.Sqrt(float64(time) / float64(rank)))
		}
		return total >= int64(cars)
	}

	for left < right {
		middle := (left + right) >> 1
		if check(int64(middle)) {
			right = middle
		} else {
			left = middle + 1
		}
	}

	return int64(left)
}
