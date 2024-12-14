package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
   "math"
)

func main() {
	var filename string = "values.txt"
	var count int = 0

	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error while opening the file: %v\n", err)
		return
	}
	defer file.Close() // Close file only if no error

	// Create a scanner to read the file line by line
	var scanner = bufio.NewScanner(file)
	
   for scanner.Scan() {
		var line = scanner.Text()
		var field = strings.Fields(line)

		// Convert input to integer array
		a := convertStringToArray(field)

		// Check conditions for validity
		if (isAscending(a) || isDescending(a)) &&  hasValidDifference(a) { 	
			   count++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading the file: %v\n", err)
	}

	fmt.Println("Total valid lines: ", count)
}

func convertStringToArray(stringArray []string) []int {
	var intArray []int
	for i := 0; i < len(stringArray); i++ {
		a, err := strconv.Atoi(stringArray[i])
		if err != nil {
			fmt.Printf("Invalid number '%s' in input. Skipping...\n", stringArray[i])
			continue
		}
		intArray = append(intArray, a)
	}
	return intArray
}

func isDescending(array []int) bool {
   for i := 0; i < len(array)-1; i++ {
		if array[i] <= array[i+1] {
         return false 
      } 
	}
	return true
}

func isAscending(array []int) bool{
   for i:= 0 ; i < len(array) - 1 ; i ++ {
      if array[i] >= array[i+1] {
         return false
      }
   }
   
   return true
}

func hasValidDifference(array []int) bool {
   for i:=0 ; i < len(array)-1 ; i++ {
      diff := float64(array[i] - array[i+1]);
      absolute := math.Abs(diff)
      if absolute < 1 || absolute > 3 {
         return false
      }
   }
   return true
}
