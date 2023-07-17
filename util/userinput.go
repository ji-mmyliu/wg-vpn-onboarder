package util

import (
	"fmt"
	"os"
)

type ValidateInputFunction[T any] func(T) bool

func GetInput[T any](prompt string, defaultVal any, validate ValidateInputFunction[T]) T {
	for true {
		var defaultValStr string
		if defaultVal == nil {
			defaultValStr = "required"
		} else {
			defaultValStr = fmt.Sprint(defaultVal)
		}

		fmt.Printf("== %s [%s]: ", prompt, defaultValStr)
		var res T
		nScanned, err := fmt.Scanln(&res)

		if defaultVal != nil && nScanned == 0 {
			def := defaultVal.(T)
			return def
		}

		if err == nil && validate(res) {
			return res
		}

		// Validation failed. Prompt for retry
		fmt.Fprintln(os.Stderr, "Invalid input; Please try again.")
	}
	panic(nil)
}
