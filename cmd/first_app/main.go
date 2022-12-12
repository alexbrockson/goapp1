package main

import (
	"fmt"
	"os"
	"strings"
	"utils/lib"
)

func main() {

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

	sum(1, 3, 4, 3, 5)
	lib.Sum(1, 3, 4, 3, 5)

}

func sum(values ...int) {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	println("Sum is", result)
}
