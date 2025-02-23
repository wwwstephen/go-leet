package main

func TwoSum(nums []int, target int) (r []int) {
	var x = map[int]int{}

	for i, num := range nums {
		if i == 0 {
			x[num] = i
			continue
		}

		compliment := target - num

		value, exists := x[compliment]

		if exists {
			r = append(r, i, value)
			return r
		} else {
			x[num] = i
		}
	}
	return r
}
