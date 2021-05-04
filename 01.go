package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	var numbers []int
	// Populate Number Array
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, line)
	}

	// Part One:
	mapnum := make(map[int]int)
	for _, i := range numbers {
		if val, ok := mapnum[i]; ok {
			fmt.Println(val * i)
		}
		mapnum[2020-i] = i
	}

	// Part Two:
	for i := 0; i < len(numbers)-2; i++ {
		target := 2020 - numbers[i]
		targetMap := make(map[int]int)
		for j := i + 1; j < len(numbers); j++ {
			if val, ok := targetMap[numbers[j]]; ok {
				fmt.Println(numbers[i] * numbers[j] * val)
			}
			targetMap[target-numbers[j]] = numbers[j]
		}
	}
}
