// conversions utilize a locally defined package to ...well, perform conversion
package main

import (
	"fmt"
	"github.com/dmithamo/learnGo/tempconv"
	"os"
	"strconv"
)

func main() {
	convertedTemp := convertTemp()
	fmt.Printf("%v\n%v\n", convertedTemp[0], convertedTemp[1])
}

// convertTemp converts temp from srcScale to destScale
func convertTemp() [2]interface{} {
	input := os.Args
	if len(input) < 3 {
		fmt.Println("Please enter a valid temp, followed by original scale and lastly, the desired scale:")

		return [2]interface{}{nil, nil}
	}

	temp, _ := strconv.ParseFloat(input[1], 64)
	srcScale, destScale := input[2], input[3]

	convertedTemp := [2]interface{}{}
	switch fmt.Sprintf("%v-%v", srcScale, destScale) {
	case "c-f":
		convertedTemp[0] = fmt.Sprintf("Temp in C: %v", tempconv.Celsius(temp))
		convertedTemp[1] = fmt.Sprintf("Temp in F: %v", tempconv.CToF(tempconv.Celsius(temp)))

	case "c-k":
		convertedTemp[0] = fmt.Sprintf("Temp in C: %v", tempconv.Celsius(temp))
		convertedTemp[1] = fmt.Sprintf("Temp in K: %v", tempconv.CToK(tempconv.Celsius(temp)))

	case "f-c":
		convertedTemp[0] = fmt.Sprintf("Temp in F: %v", tempconv.Fahrenheit(temp))
		convertedTemp[1] = fmt.Sprintf("Temp in C: %v", tempconv.FToC(tempconv.Fahrenheit(temp)))

	case "f-k":
		convertedTemp[0] = fmt.Sprintf("Temp in F: %v", tempconv.Fahrenheit(temp))
		convertedTemp[1] = fmt.Sprintf("Temp in K: %v", tempconv.FToK(tempconv.Fahrenheit(temp)))
	}

	return convertedTemp
}
