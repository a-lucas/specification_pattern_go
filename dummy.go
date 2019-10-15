package specification_pattern

import (
	"log"
	"strconv"
)

/* This file's purpose is to hack & hide the original code

 */

func PointToPeriod(nbPoints int) string {
	return strconv.Itoa(nbPoints) + " points"
}

func FloatToStringPrecision(input_num float64, precision int) string {
	// to convert a float number to a string
	return strconv.FormatFloat(input_num, 'f', precision, 64)
}

func StrToFloat(input string) float64 {
	val, err := strconv.ParseFloat(input, 64)
	if err != nil {
		log.Println("Error parsing string to float", input)
		log.Panicln(err)
	}
	return val
}
