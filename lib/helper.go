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
