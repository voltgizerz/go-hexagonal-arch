package utils

import (
	"fmt"
	"strconv"
)

func ToString[T any](value T) string {
	return fmt.Sprintf("%v", value)
}

func StringToInt(value string) int {
	num, err := strconv.Atoi(value)
	if err != nil {
		return 0
	}
	return num
}
