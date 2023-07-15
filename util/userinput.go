package util

import "fmt"

type ValidateInputFunction[T any] func(T) bool

func GetInput[T any](prompt string, def T, validate ValidateInputFunction[T]) T {
	for true {
		fmt.Printf("%s [%s]: ", prompt, fmt.Sprint(def))
		var res T
		nScanned, err := fmt.Scanln(&res)

		fmt.Println("Scanned:", nScanned)
		if nScanned == 0 {
			return def
		}

		if err == nil && validate(res) {
			return res
		}

		// Validation failed. Prompt for retry
		fmt.Println("Invalid input; Please try again.")
	}
	panic(nil)
}
