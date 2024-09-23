package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go data.txt")
	} else {

		filename := os.Args[1]
		data, err := readData(filename)
		if err != nil {
			fmt.Printf("Error reading data: %v\n", err)
			return
		}
		if len(data) == 0 {
			fmt.Println("Error: file is empty")
			return
		}

		n := float64(len(data))
		sumX, sumY, sumXY, sumX2, sumY2 := 0.0, 0.0, 0.0, 0.0, 0.0

		for i, y := range data {
			x := float64(i)
			sumX += x
			sumY += y
			sumXY += x * y
			sumX2 += x * x
			sumY2 += y * y
		}

		// Calculate Linear Regression coefficients
		slope := (n*sumXY - sumX*sumY) / (n*sumX2 - sumX*sumX)
		intercept := (sumY - slope*sumX) / n

		// Calculate Pearson Correlation Coefficient
		numerator := n*sumXY - sumX*sumY
		denominator := math.Sqrt((n*sumX2 - sumX*sumX) * (n*sumY2 - sumY*sumY))
		pearson := numerator / denominator

		// Print results with required precision
		fmt.Printf("Linear Regression Line: y = %.6fx + %.6f\n", slope, intercept)
		fmt.Printf("Pearson Correlation Coefficient: %.10f\n", pearson)
	}
}

func readData(filename string) ([]float64, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("could not open file: %v", err)
	}
	defer file.Close()

	var data []float64
	scanner := bufio.NewScanner(file)
	count := 0
	for scanner.Scan() {
		count++
		line := scanner.Text()
		value, err := strconv.ParseFloat(line, 64)
		if err != nil {
			fmt.Println(fmt.Errorf("invalid number in file line %d : %v ", count, err))
		}
		data = append(data, value)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	return data, nil
}
