package main

import (
	"fmt"
	"unicode"
	"errors"
)

func parse(str string) (int, error) {
	i := 0

    for _, c := range str {
		if !unicode.IsDigit(c) {
			return  -1, errors.New("invalid int")
		}
		digit := int(c) - '0'
		i = i * 10 + digit
	}

	return i, nil
}

func parseFloat(str string) (float64, error) {
	i := 0.0
	decimal := 0.1
	float_pointer := false

    for _, c := range str {
		if !unicode.IsDigit(c) && c != '.' && float_pointer {
			return  -1, errors.New("invalid float")
		}
		if c == '.' {
			float_pointer = true
		} else if float_pointer {
			digit := int(c) - '0'
			i = i + (decimal * float64(digit))
			decimal = decimal * 0.1
		} else {
			digit := int(c) - '0'
		    i = i * 10 + float64(digit)
		}
	}

	return i, nil
}

func main() {
	fmt.Println(parse("23300"))
	fmt.Println(parseFloat("233.0234"))
}
