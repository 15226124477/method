package method

import (
	"encoding/binary"
	"strconv"
	"strings"
)

// PadLeft 字符串左补0
func PadLeft(str string, length int) string {
	for len(str) < length {
		str = "0" + str
	}
	return str
}

// XorChecksum 异或校验
func XorChecksum(str string) string {
	var checksum byte
	for i := 0; i < len(str); i++ {
		checksum ^= str[i]
	}
	result := int(binary.BigEndian.Uint16([]byte{0, checksum}))
	FF := strings.ToUpper(strconv.FormatInt(int64(result), 16))
	if len(FF) == 1 {
		return "0" + FF
	} else {
		return FF
	}
}

// RemoveRepeatedElement 数组去重
func RemoveRepeatedElement(items interface{}) interface{} {
	switch items.(type) {
	case []int:
		intArr := items.([]int)
		result := make([]int, 0)
		m := make(map[int]bool) //map的值不重要
		for _, v := range intArr {
			if _, ok := m[v]; !ok {
				result = append(result, v)
				m[v] = true
			}
		}
		return result
	case []string:
		strArr := items.([]string)
		result := make([]string, 0)
		m := make(map[string]bool) //map的值不重要
		for _, v := range strArr {
			if _, ok := m[v]; !ok {
				result = append(result, v)
				m[v] = true
			}
		}
		return result
	case []float64:
		strArr := items.([]float64)
		result := make([]float64, 0)
		m := make(map[float64]bool) //map的值不重要
		for _, v := range strArr {
			if _, ok := m[v]; !ok {
				result = append(result, v)
				m[v] = true
			}
		}
		return result
	default:
		return nil
	}

}

func IsContain(items interface{}, item interface{}) bool {
	switch items.(type) {
	case []int:
		intArr := items.([]int)
		for _, value := range intArr {
			if value == item.(int) {
				return true
			}
		}
	case []string:
		strArr := items.([]string)
		for _, value := range strArr {
			if value == item.(string) {
				return true
			}
		}
	case []float64:
		strArr := items.([]float64)
		for _, value := range strArr {
			if value == item.(float64) {
				return true
			}
		}
	default:
		return false
	}
	return false
}

// Average 浮点数组平均值
func Average(xs interface{}) (avg float64) {
	switch xs.(type) {
	case []int:
		items := xs.([]int)
		sum := 0.0
		switch len(items) {
		case 0:
			avg = 0
		default:
			for _, v := range items {
				sum += float64(v)
			}
			avg = sum / float64(len(items))
		}
		return avg
	case []float64:
		items := xs.([]float64)
		sum := 0.0
		switch len(items) {
		case 0:
			avg = 0
		default:
			for _, v := range items {
				sum += v
			}
			avg = sum / float64(len(items))
		}
		return avg
	default:
		return 0
	}

}

// ListCount 列表计数
func ListCount(List interface{}) interface{} {
	switch List.(type) {
	case []int:
		hash := make(map[int]int)
		for _, item := range List.([]int) {
			if _, ok := hash[item]; ok {
				hash[item]++
			} else {
				hash[item] = 1
			}
		}
		return hash
	case []float64:
		hash := make(map[float64]int)
		for _, item := range List.([]float64) {
			if _, ok := hash[item]; ok {
				hash[item]++
			} else {
				hash[item] = 1
			}
		}
		return hash
	default:
		return 0
	}

}
