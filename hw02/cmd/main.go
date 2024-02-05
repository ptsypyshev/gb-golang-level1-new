package main

import (
	"fmt"
	"log"
	"os"
	"unicode"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Укажите предложение в кавычках вторым аргументом. Например, ./main \"hello world\"\n")
	}

	phrase := os.Args[1]

	countMap, total := getCountMapAndTotalCount(phrase)
	for _, k := range phrase {
		if v, ok := countMap[k]; ok {
			fmt.Printf("%c - %.0f %.2f\n", k, v, v/total)
			delete(countMap, k)
		}
	}
}

func getCountMapAndTotalCount(phrase string) (map[rune]float64, float64) {
	var total float64
	countMap := make(map[rune]float64)
	for _, char := range phrase {
		if unicode.IsLetter(char) {
			countMap[char]++
			total++
		}
	}
	return countMap, total
}
