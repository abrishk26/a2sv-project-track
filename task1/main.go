package main

import (
	"fmt"
	// "bufio"
	// "os"
	"strings"
	"strconv"
)

func main() {
	// reader := bufio.NewReader(os.Stdin)

	
	fmt.Println("Hello World!")
}

func parseGrades(grades string) ([]float64, error) {
	res := []float64{}

	for _, field := range strings.Fields(grades) {
		grade, err := strconv.ParseFloat(field, 64)
		if err != nil {
			return nil, err 
		}

		res = append(res, grade)
	}

	return res, nil
}
