package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func formatPrintNums(nums string) {
	var nums_slice = strings.Split(nums, "\n")
	var count = 0

	for i := 0; i < len(nums_slice); i++ {
		var inner_slice = strings.Split(nums_slice[i], "")

		for j := 0; j < len(inner_slice); j++ {
			if inner_slice[j] == "5" {
				count += 1
				fmt.Printf("%v -- %v\n", count, nums_slice[i])
			}
		}
	}

	fmt.Printf("Through 1-100, the number 5 comes %v times\n", count)
}

func main() {
	var data, err = os.ReadFile("numbers.txt")

	if err != nil {
		log.Fatal(err)
	}

	var nums = string(data)

	formatPrintNums(nums)
}
