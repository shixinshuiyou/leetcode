package middle

// https://leetcode.cn/problems/median-of-two-sorted-arrays/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	length := len(nums1) + len(nums2)
	mid1, mid2 := length/2, length/2
	if length%2 == 0 {
		mid1 = length/2 - 1
	}
	//fmt.Println(mid1, mid2)
	i, j := 0, 0
	var nums []int
	var ver int
	for true {
		if i >= len(nums1) {
			nums = append(nums, nums2[j])
			j++
		} else if j >= len(nums2) {
			nums = append(nums, nums1[i])
			i++
		} else if nums1[i] < nums2[j] {
			nums = append(nums, nums1[i])
			i++
		} else {
			nums = append(nums, nums2[j])
			j++
		}
		if i+j-1 == mid1 {
			ver += nums[i+j-1]
			//fmt.Println(nums[i+j-1])
		}
		if i+j-1 == mid2 {
			ver += nums[i+j-1]
			//fmt.Println(nums[i+j-1])
			break
		}
	}
	return float64(ver) / 2
}

// https://leetcode.cn/problems/regular-expression-matching/
func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	matches := func(i, j int) bool {
		if i == 0 {
			return false
		}
		if p[j-1] == '.' {
			return true
		}
		return s[i-1] == p[j-1]
	}

	f := make([][]bool, m+1)
	for i := 0; i < len(f); i++ {
		f[i] = make([]bool, n+1)
	}
	f[0][0] = true
	for i := 0; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '*' {
				f[i][j] = f[i][j] || f[i][j-2]
				if matches(i, j-1) {
					f[i][j] = f[i][j] || f[i-1][j]
				}
			} else if matches(i, j) {
				f[i][j] = f[i][j] || f[i-1][j-1]
			}
		}
	}
	return f[m][n]
}
