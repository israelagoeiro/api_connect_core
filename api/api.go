package api

import (
	"fmt"
	"time"
)

type Error struct {
	date    string
	message string
}

func (e Error) Error() string {
	return fmt.Sprintf("%v: %v", e.date, e.message)
}

func PanicIsEmpty(value string, source string, message string) {
	if !(len(value) > 0) {
		panic(Error{
			time.Now().Format("2006-01-02 15:04:05"),
			source + "-> " + message,
		})
	}
}
