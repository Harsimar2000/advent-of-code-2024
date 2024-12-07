package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
  "sort"
  "math"
)

func main() {
	// Specify the file name
	fileName := "values.txt"

	// Open the file
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Slices to hold the numbers for each column
	var column1 []int
	var column2 []int

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Read the current line
		line := scanner.Text()

		// Split the line into columns (assume space-separated values)
		fields := strings.Fields(line)

		// Ensure there are at least two numbers per line
		if len(fields) >= 2 {
			// Convert and append the first number to column1
			num1, err := strconv.Atoi(fields[0])
			if err != nil {
				fmt.Println("Error converting to integer:", err)
				continue
			}
			column1 = append(column1, num1)

			// Convert and append the second number to column2
			num2, err := strconv.Atoi(fields[1])
			if err != nil {
				fmt.Println("Error converting to integer:", err)
				continue
			}
			column2 = append(column2, num2)
		}
	}

	// Check for scanner errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
  
  // Sort the slices
  sort.Ints(column1)
  sort.Ints(column2)

  var sum_of_difference int = 0;
 for i := 0 ; i < len(column1) ; i++ {
  sum_of_difference += int(math.Abs(float64(column1[i] - column2[i])))
 }
 
 fmt.Println(sum_of_difference)
	// Print the resulting slices?..
	fmt.Println("Column 1:", len(column1))
	fmt.Println("Column 2:", len(column2))
}

