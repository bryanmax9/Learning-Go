package conversion

import (
	"errors"
	"strconv"
)

// Coverting a list of strings into a list of float64
func StringsToFloats(lines []string) ([]float64, error) {
	var floats []float64

	for _, line := range lines {
		floatPrice, err := strconv.ParseFloat(line, 64)

		if err != nil {

			return nil, errors.New("Failed to convert string to float")
		}

		floats = append(floats, floatPrice)
	}

	return floats, nil
}
