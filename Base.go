package method

func RemoveRepeatedElement(s []string) []string {
	result := make([]string, 0)
	m := make(map[string]bool) //map的值不重要
	for _, v := range s {
		if _, ok := m[v]; !ok {
			result = append(result, v)
			m[v] = true
		}
	}
	return result
}

func IsContain(items []string, item string) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}

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
	for _, item := range floatList {
		if _, ok := hash[item]; ok {
			hash[item]++
		} else {
			hash[item] = 1
		}
	}
	return hash
}
