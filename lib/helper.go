package lib

import (
	"fmt"
)

func Sum(values ...int) {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	println("Sum is", result)
}
