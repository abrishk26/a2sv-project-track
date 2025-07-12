package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
	"errors"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your Full Name: ")
	fullName, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Invalid input")
		return 
	}

	fullName = strings.Trim(fullName, "\n")

	fmt.Print("Enter number of courses you have taken: ")
	
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Invalid input")
		return 
	}

	courseCount, err := strconv.ParseInt(strings.Trim(input, "\n"), 10, 64)
	if err != nil {
		fmt.Println("Invalid number of courses. Please try again.")
		return
	}

	fmt.Print("Enter course names: ")
	names, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Invalid input")
		return
	}

	courseNames := strings.Fields(names)
	if int64(len(courseNames)) != courseCount {
		fmt.Printf("Expected %d courses, found %d courses", courseCount, len(courseNames))
		return
	}

	fmt.Print("Enter course grades: ")
	grades, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Invalid input")
		return 
	}

	courseGrades, err := parseGrades(grades)
	if err != nil {
		fmt.Println("Invalid grades")
		return
	}

	
	if int64(len(courseGrades)) != courseCount {
		fmt.Printf("Expected %d courses, found %d courses", courseCount, len(courseGrades))
		return
	}
	

	
	fmt.Printf("Full Name: %s\nNumber of Courses: %d\n", fullName, courseCount)
}

func parseGrades(grades string) ([]float64, error) {
	res := []float64{}

	for _, field := range strings.Fields(grades) {
		grade, err := strconv.ParseFloat(field, 64)
		if err != nil {
			return nil, err 
		}

		if grade < 0 || grade > 100 {
			return nil, errors.New("Invalid grade value") 
		}

		res = append(res, grade)
	}

	return res, nil
}

func calculateAvg(grades []float64) float64 {
	var sum float64

	for _, grade := range grades {
		sum += grade
	}

	return sum / float64(len(grades))
}
