package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	file, err := os.Open("./input/day_1.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0
	for scanner.Scan() {
		for _, c := range scanner.Text() {
			if unicode.IsDigit(c) {
				sum += int(c)
				fmt.Println(c)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(sum)
}
