package main

import (
	"fmt"
	"os"
	"strings"
	// "goapp1/test"
	// "controllers/test"
)

func main() {

	// var j int = 3434
	// fmt.Printf("%v, %T\n", j, j)

	//for i := 0; i < 5; i++ {
	//	fmt.Println(i)
	//	fmt.Printf("%v, %T\n", i, i)
	//}

	// fmt.Println(count)

	f, err := os.ReadFile("data/data.txt")
	splitStrings := strings.Split(string(f), "\n")
	// formattedString, test, test2 := strings.Cut(string(f), "\n")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%v%v\n", "Number of words: ", len(splitStrings))
	for i := 0; i < len(splitStrings); i++ {
		fmt.Printf("%v\n", splitStrings[i])
		// fmt.Println(string(test))
		// fmt.Println(test2)
	}

}
