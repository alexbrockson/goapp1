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

}
