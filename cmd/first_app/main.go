package main

import (
	"goapp1/utils/lib"
	"fmt"
	"os"
	"strings"
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

}
