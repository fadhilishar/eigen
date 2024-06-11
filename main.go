// Algoritma
package main

import (
	"fmt"
	"strconv"
)

func main() {

	// 1
	word := "NEGIE1"
	fmt.Println(refWord(word))

	// 2
	sentence := "Saya sangat senang mengerjakan soal algoritma"
	fmt.Println(longestWord(sentence))

	// 3
	input := []string{"xc", "dz", "bbb", "dz"}
	query := []string{"bbb", "ac", "dz"}
	outputCount := wordCount(query, input)
	fmt.Println("outputCount: ", outputCount)

	// 4
	matrix := [][]int{{1, 2, 0}, {4, 5, 6}, {7, 8, 9}}
	fmt.Printf("selisih diagonalnya adalah %v", diffDiagonalSumMatrix(matrix))
}

// 1
func refWord(word string) (refWord string) {
	numWord := ""
	for i := len(word) - 1; i >= 0; i-- {
		if isNum(string(word[i])) {
			numWord = string(word[i]) + numWord
			continue
		}

		refWord += string(word[i])
	}
	refWord += numWord
	return
}

func isNum(word string) bool {
	if _, err := strconv.Atoi(word); err != nil {
		return false
	}
	return true
}

// 2
func longestWord(sentence string) string {
	word := ""
	temp := ""
	// countTemp := 0
	// countMax := 0
	// maxLength := 0
	for i := 0; i < len(sentence); i++ {
		// fmt.Printf()
		if string(sentence[i]) != " " {
			// countTemp++
			temp += string(sentence[i])
			continue
		}
		// if countTemp > countMax {
		// 	countMax = countTemp
		// }

		// else {
		if len(temp) > len(word) {
			word = temp
		}
		temp = ""
		// }

		// temp = ""
	}
	return fmt.Sprintf("%v: %v character", word, len(word))
}

// 3
func wordCount(arrQuery []string, arrInput []string) (arrOutput []int) {

	// mapWordToIsExist := map[string]bool{}
	mapWordToCount := map[string]int{}
	// arrOutput = make([]int, len(arrQuery))
	for _, wordInput := range arrInput {
		mapWordToCount[wordInput]++
	}

	for _, wordQuery := range arrQuery {
		arrOutput = append(arrOutput, mapWordToCount[wordQuery])
		// arrOutput[i]
	}

	return
}

// 4
func diffDiagonalSumMatrix(matrix [][]int) int {
	// ki-ka, ka-ki
	diagonal1, diagonal2 := 0, 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if i == j {
				diagonal1 += matrix[i][j]
			}
			if i+j == len(matrix)-1 {
				diagonal2 += matrix[i][j]
			}
		}
	}
	if diagonal1 >= diagonal2 {
		return diagonal1 - diagonal2
	}

	return diagonal2 - diagonal1
}
