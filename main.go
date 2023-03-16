package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

var originUnit string
var originValue float64

var shouldConvertAgain string

var err error

var errInvalidArguments = errors.New("Invalid arguments")
var errReadingInput = errors.New("Error reading input")

func main() {
	if len(os.Args) != 2 {
		log.Fatalln(errInvalidArguments.Error())
	}

	originUnit = strings.ToUpper(os.Args[1])
	var input float64

	for {
		fmt.Print("What is the current temperature in " + originUnit + " ? ")

		switch originUnit {
		case "C":
			n, err := fmt.Scanf("%f", &input)
			if n == 0 || err != nil {
				log.Fatalln(errReadingInput)
			}
			convertToFahrenheit(input)
		case "F":
			n, err := fmt.Scanf("%f", &input)
			if n == 0 || err != nil {
				log.Fatalln(errReadingInput)
			}
			convertToCelsius(input)
		default:
			log.Fatalln("invalid Unit: ", originUnit)
		}

		fmt.Print("Would you like to convert another temperature ? (y/n) ")

		fmt.Scanln(&shouldConvertAgain)

		shouldConvertAgain = strings.ToUpper(shouldConvertAgain)

		if shouldConvertAgain != "Y" {
			fmt.Println("Good bye!")
			break
		}
	}
}

func printError(err error) {
	fmt.Fprintf(os.Stderr, "error: %v\n", err)
	os.Exit(1)
}

func convertToCelsius(value float64) {
	convertedValue := (value - 32) * 5 / 9
	fmt.Printf("%v F = %.0f C\n", value, convertedValue)
}

func convertToFahrenheit(value float64) {
	convertedValue := (value * 9 / 5) + 32
	fmt.Printf("%v C = %.0f F\n", value, convertedValue)
}
