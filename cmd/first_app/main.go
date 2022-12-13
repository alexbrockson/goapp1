package main

import (
	"fmt"
	"goapp1/utils/lib"
	"os"
	"strconv"
	"strings"
)

func main() {
	defer println() // adds a new line to the end of the output

	f, err := os.ReadFile("data/data.txt")
	splitStrings := strings.Split(string(f), "\n")

	if err != nil {
		fmt.Println(err)
		println(err)
		return
	}
	names := make([]string, 0)
	values := make([]int, 0)

	println("Number of words:", len(splitStrings))
	for i := 0; i < len(splitStrings); i++ {
		line := strings.Split(splitStrings[i], "=")
		name := strings.TrimSpace(line[0])
		names = append(names, name)
		nums, err := strconv.Atoi(strings.TrimSpace(line[1]))
		if err == nil {
			values = append(values, nums)
		}
	}
	fmt.Println(names, values)
	test := lib.Exclamation
	println("This:", test, "WHAT'S UP")

	//sum(1, 3, 4, 3, 5)
	println("\nSum function")
	println(lib.Sum(1, 3, 4, 3, 5))

	println("\nDivide function")
	d, err := lib.Divide(128.0, 8)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)

	println("\nDivide function (zero test)")
	i, err2 := lib.Divide(3.0, 0.0)
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	fmt.Println(i)

}

// func sum(values ...int) {
// 	fmt.Println(values)
// 	result := 0
// 	for _, v := range values {
// 		result += v
// 	}
// 	println("Sum is", result)
// }
