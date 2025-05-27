package utils

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func ReadLine(scanner *bufio.Scanner) string {
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

func ReadFloat(scanner *bufio.Scanner) float64 {
	text := ReadLine(scanner)
	val, err := strconv.ParseFloat(text, 64)
	if err != nil {
		fmt.Println("Invalid number.")
		return 0
	}
	return val
}
