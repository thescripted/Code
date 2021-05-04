package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
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

	scanner := bufio.NewScanner(file)
	counterOne := 0
	counterTwo := 0
	for scanner.Scan() {
		if checkPasswordCount(scanner.Text()) {
			counterOne += 1
		}
		if checkValidPassword(scanner.Text()) {
			counterTwo += 1
		}
	}
	fmt.Printf("RESULTS:\n\nCounter One: %d\nCounter Two:%d\n", counterOne, counterTwo)
}

func checkPasswordCount(input string) bool {
	leader, trailer, token, data := parseInput(input)
	count := 0
	for _, char := range data {
		if string(char) == token {
			count += 1
		}
	}

	if leader <= count && trailer >= count {
		return true
	}
	return false
}

func checkValidPassword(input string) bool {
	leader, trailer, token, data := parseInput(input)
	if (token == string(data[leader-1])) != (token == string(data[trailer-1])) {
		return true
	}
	return false
}

func parseInput(input string) (int, int, string, string) {
	r := regexp.MustCompile(`^(\d+)-(\d+)\s([A-Za-z]):\s(\S+)$`)
	parsed := r.FindStringSubmatch(input)
	leader, _ := strconv.Atoi(parsed[1])
	trailer, _ := strconv.Atoi(parsed[2])
	token := parsed[3]
	data := parsed[4]
	return leader, trailer, token, data
}
