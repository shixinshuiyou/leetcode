package middle

func rob(nums []int) int {
	k, j := 0, 0
	for _, n := range nums {
		k, j = j, max(k+n, j)
	}
	return max(k, j)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
