package logger

import "fmt"

func E(err error) {
	fmt.Printf("[E] %v\n", err.Error())
}

func W(err error) {
	fmt.Printf("[E] %v\n", err.Error())
}
