package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func readNumber(text string) (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(text)
		input, err := reader.ReadString('\n')
		if err != nil {
			return 0, fmt.Errorf("incorrect enter, error: %v", err)
		}
		input = strings.TrimSpace(input)
		value, err := strconv.ParseFloat(input, 64)
		if err != nil {
			return 0, fmt.Errorf("incorrect enter, error: %v, please try again", err)
		}
		return value, nil
	}
}

func getUserInput() (float64, float64, error) {
	height, err := readNumber("Enter your height: ")
	if err != nil {
		return 0, 0, err
	}
	weight, err := readNumber("Enter your weight: ")
	if err != nil {
		return 0, 0, err
	}

	if height <= 0 {
		return 0, 0, fmt.Errorf("height must be greater than zero")
	}
	if weight <= 0 {
		return 0, 0, fmt.Errorf("weight must be greater than zero")
	}
	return height, weight, nil
}

func calculateIMT(height float64, weight float64) float64 {
	return weight / math.Pow(height/100, 2)
}

func outputResult(imt float64) {
	fmt.Printf("Your IMT: %.2f\n", imt)
	switch {
	case imt < 16:
		fmt.Println("Severe body weight deficiency")
	case imt < 18.5:
		fmt.Println("Body weight deficiency")
	case imt < 25:
		fmt.Println("Body weight is normal")
	case imt < 30:
		fmt.Println("Excess weight")
	case imt < 35:
		fmt.Println("The first degree of obesity")
	case imt < 40:
		fmt.Println("The second degree of obesity")
	default:
		fmt.Println("The third degree of obesity")
	}
}

func checkRepeatCalculation() bool {
	var userInput string
	fmt.Print("Do you want to repeat? (Yes/No): ")
	fmt.Scan(&userInput)
	if userInput == "Yes" || userInput == "yes" {
		return true
	}
	return false
}

func main() {
	fmt.Println("Hi! This is IMT calculator.")
	for {
		height, weight, err := getUserInput()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		IMT := calculateIMT(height, weight)
		outputResult(IMT)
		isRepeat := checkRepeatCalculation()
		if isRepeat {
			fmt.Println("Let's start the calculation from the beginning...")
		} else {
			fmt.Println("Thanks for using the calculation!")
			break
		}
	}
}
