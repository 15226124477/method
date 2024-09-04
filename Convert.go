package method

import (
	"math"
	"strconv"
)

const (
	DEGREES = 1
	RADIAN  = 2

	// GGASOL = 3
	// POSSOL = 4
)

// Decimal 浮点数保留N位小数
func Decimal(value float64, prec int) float64 {
	value, _ = strconv.ParseFloat(strconv.FormatFloat(value, 'f', prec, 64), 64)
	return value
}

// Degrees2Radians 度转弧度
func Degrees2Radians(degrees float64) float64 {
	return degrees * math.Pi / 180
}

// Radians2Degrees 弧度转度
func Radians2Degrees(radians float64) float64 {
	return radians * 180 / math.Pi
}
