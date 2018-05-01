package errors

import (
	"fmt"
	"log"
)

func FailOnError(message string, err error) {
	if err != nil {
		log.Printf("%s. Reason: %s", message, err)
		panic(fmt.Sprintf("%s. Reason: %s", message, err))
	}
}

func LogOnError(message string, err error) {
	if err != nil {
		log.Printf("Oops! %s. Error: %s", message, err)
	}
}
