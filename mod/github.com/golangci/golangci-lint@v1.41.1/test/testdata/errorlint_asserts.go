//args: -Eerrorlint
//config_path: testdata/configs/errorlint_asserts.yml
package testdata

import (
	"errors"
	"log"
)

type myError struct{}

func (*myError) Error() string {
	return "foo"
}

func errorLintDoAnotherThing() error {
	return &myError{}
}

func errorLintAsserts() {
	err := errorLintDoAnotherThing()
	var me *myError
	if errors.As(err, &me) {
		log.Println("myError")
	}
	_, ok := err.(*myError) // ERROR "type assertion on error will fail on wrapped errors. Use errors.As to check for specific errors"
	if ok {
		log.Println("myError")
	}
	switch err.(type) { // ERROR "type switch on error will fail on wrapped errors. Use errors.As to check for specific errors"
	case *myError:
		log.Println("myError")
	}
	switch errorLintDoAnotherThing().(type) { // ERROR "type switch on error will fail on wrapped errors. Use errors.As to check for specific errors"
	case *myError:
		log.Println("myError")
	}
	switch t := err.(type) { // ERROR "type switch on error will fail on wrapped errors. Use errors.As to check for specific errors"
	case *myError:
		log.Println("myError", t)
	}
	switch t := errorLintDoAnotherThing().(type) { // ERROR "type switch on error will fail on wrapped errors. Use errors.As to check for specific errors"
	case *myError:
		log.Println("myError", t)
	}
}
