package method

// Average 浮点数组平均值
func Average(xs []float64) (avg float64) {
	sum := 0.0
	switch len(xs) {
	case 0:
		avg = 0
	default:
		for _, v := range xs {
			sum += v
		}
		avg = sum / float64(len(xs))
	}
	return
}

// ListCount 列表计数
func ListCount(floatList []float64) map[float64]int {
	hash := make(map[float64]int)
	// 遍历数组，将每个元素插入map中
	// 如果元素已经存在，计数器加一；如果不存在，则插入并将计数器置为1
	for _, item := range floatList {
		if _, ok := hash[item]; ok {
			hash[item]++
		} else {
			hash[item] = 1
		}
	}
	return hash
}

// IsStringContain 字符串是否在数组中
func IsStringContain(items []string, item string) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}
