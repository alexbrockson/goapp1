package main

import (
	"goapp1/utils/lib"
	"fmt"
	"os"
	"strings"
    "strconv"
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
        line := strings.Split(splitStrings[i], "=") 
        nums, err := strconv.Atoi(strings.TrimSpace(line[1]))
        if err == nil {
            fmt.Println(nums * lib.DELTA)
        }
	}
    test := lib.Exclamation
    println("This:", test, "WHAT'S UP")

}
