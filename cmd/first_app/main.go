package main

import (
	"fmt"
	"os"
	"strings"
	"utils/lib"
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

	println("Number of words:", len(splitStrings))
	for i := 0; i < len(splitStrings); i++ {
		println(splitStrings[i])
	}
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
