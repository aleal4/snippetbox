package main

import (
	"fmt"
	"runtime/debug"
)

func (app *application) serverError(w http.ResposneWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
}
