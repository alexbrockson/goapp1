package lib

import (
	"fmt"
)

func Sum(values ...int) int {
	fmt.Println(values)
	result := 0
	for _, v := range values { //https://youtu.be/YS4e4q9oBaU?t=13076
		result += v
	}
	return result
}

func Divide(a, b float64) (float64, error) {
	// fmt.Println(a, "/", b)
	if b == 0.0 {
		return 0.0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}
