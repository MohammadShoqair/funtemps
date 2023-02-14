package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	// Define flags for input temperature value, input temperature unit,
	// and output temperature unit
	var temp float64
	var inUnit, outUnit string
	fmt.Println("Usage: conv <value> <input unit> <output unit>")
	fmt.Println("Example: conv 32 F C")
	fmt.Println()

	if len(os.Args) < 4 {
		fmt.Println("Error: insufficient arguments")
		os.Exit(1)
	}

	temp, inUnit, outUnit = parseArgs(os.Args[1:])
	result := convertTemp(temp, inUnit, outUnit)
	fmt.Printf("%.2f%s er %.2f%s\n", temp, inUnit, result, outUnit)
}

func parseArgs(args []string) (float64, string, string) {
	temp, err := parseTemp(args[0])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	inUnit := strings.ToUpper(args[1])
	outUnit := strings.ToUpper(args[2])
	if inUnit == outUnit {
		fmt.Println("Error: input and output units are the same")
		os.Exit(1)
	}

	return temp, inUnit, outUnit
}

func parseTemp(arg string) (float64, error) {
	var temp float64
	_, err := fmt.Sscanf(arg, "%f", &temp)
	if err != nil {
		return 0, fmt.Errorf("invalid temperature value: %s", arg)
	}

	return temp, nil
}

func convertTemp(temp float64, inUnit string, outUnit string) float64 {
	if inUnit == "C" && outUnit == "F" {
		return temp*9/5 + 32
	} else if inUnit == "C" && outUnit == "K" {
		return temp + 273.15
	} else if inUnit == "F" && outUnit == "C" {
		return (temp - 32) * 5 / 9
	} else if inUnit == "F" && outUnit == "K" {
		return (temp + 459.67) * 5 / 9
	} else if inUnit == "K" && outUnit == "C" {
		return temp - 273.15
	} else if inUnit == "K" && outUnit == "F" {
		return temp*9/5 - 459.67
	} else if inUnit == "C" && outUnit == "K" {
		return temp + 273.15
	} else if inUnit == "K" && outUnit == "C" {
		return temp - 273.15
	} else {
		fmt.Printf("Error: invalid unit combination: %s to %s\n", inUnit, outUnit)
		os.Exit(1)
	}

	return math.NaN()
}

func GetSunTemp() string {
	return "Temperatur på ytre lag av Solen 5506.85°C.\nTemperatur i Solens kjerne er 15 000 000°C."
}
