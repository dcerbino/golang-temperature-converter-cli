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

var errInvalidArguments = errors.New("invalid arguments")
var errReadingInput = errors.New("error reading input")

func main() {
	if len(os.Args) != 2 {
		log.Fatalln(errInvalidArguments.Error())
	}

	originUnit = strings.ToUpper(os.Args[1])
	var inputs int

	for {
		fmt.Print("What is the current temperature in " + originUnit + " ? ")

		switch originUnit {
		case "C":
			inputs, err = fmt.Scanf("%f", &originValue)
			if inputs == 0 || err != nil {
				printError(errReadingInput)
			}
			convertToFahrenheit(originValue)
		case "F":
			inputs, err = fmt.Scanf("%f", &originValue)
			if inputs == 0 || err != nil {
				printError(errReadingInput)
			}
			convertToCelsius(originValue)
		default:
			printError(errors.New("invalid Origin unit"))
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
