package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func Execute() {
	// open file
	fh, err := os.Open("day1/input")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer fh.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(fh)

	var allCalories []int
	maxCalories, currentCalories := 0, 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if currentCalories > maxCalories {
				maxCalories = currentCalories
			}
			allCalories = append(allCalories, currentCalories)
			currentCalories = 0
			continue
		}

		calories, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		currentCalories += calories
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	
	sort.Slice(allCalories, func(i, j int) bool {
		return allCalories[i] > allCalories[j]
	})

	sum := allCalories[0] + allCalories[1] + allCalories[2]
	fmt.Println(sum)
}
