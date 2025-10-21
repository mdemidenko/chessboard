package validator

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const maxSize = 100

func ValidateSize(input string) (int, error) {
	input = strings.TrimSpace(input)

	if input == "" {
		return 0, errors.New("размер не может быть пустым")
	}

	size, err := strconv.Atoi(input)
	if err != nil {
		return 0, errors.New("размер должен быть целым числом")
	}

	if size <= 0 {
		return 0, errors.New("размер должен быть положительным числом")
	}

	if size > maxSize {
		return 0, fmt.Errorf("размер не может превышать %d", maxSize)
	}

	return size, nil
}
