package conversion

import (
	"errors"
	"fmt"
	"strconv"
)

func StringsToFloat(strings []string) ([]float64, error) {
	floatValues := make([]float64, len(strings))
	for strIdx, str := range strings {
		value, err := strconv.ParseFloat(str, 64)

		if err != nil {
			fmt.Println("Error converting string to float64")
			fmt.Println(err)
			return nil, errors.New("Error converting string to float64")
		}
		floatValues[strIdx] = value
	}

	return floatValues, nil
}
