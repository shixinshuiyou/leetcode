package middle

// 20. 有效的括号
func isValid(s string) bool {
	var data []int32
	var lastNode, size int32

	check := func(s1, s2 int32) bool {
		if s1 == '(' && s2 == ')' {
			return true
		}
		if s1 == '[' && s2 == ']' {
			return true
		}
		if s1 == '{' && s2 == '}' {
			return true
		}
		return false
	}

	for _, b := range s {
		if size == 0 || !check(lastNode, b) {
			data = append(data, b)
			size++
			lastNode = b
			continue
		}
		size--
		if size == 0 {
			data = make([]int32, 0)
		} else {
			data = data[:size]
			lastNode = data[size-1]
		}

	}
	return size == 0
}
