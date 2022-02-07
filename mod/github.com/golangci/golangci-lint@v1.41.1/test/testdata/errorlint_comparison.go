//args: -Eerrorlint
//config_path: testdata/configs/errorlint_comparison.yml
package testdata

import (
	"errors"
	"log"
)

var errCompare = errors.New("foo")

func errorLintDoThing() error {
	return errCompare
}

func errorLintComparison() {
	err := errorLintDoThing()
	if errors.Is(err, errCompare) {
		log.Println("errCompare")
	}
	if err == nil {
		log.Println("nil")
	}
	if err != nil {
		log.Println("nil")
	}
	if nil == err {
		log.Println("nil")
	}
	if nil != err {
		log.Println("nil")
	}
	if err == errCompare { // ERROR "comparing with == will fail on wrapped errors. Use errors.Is to check for a specific error"
		log.Println("errCompare")
	}
	if err != errCompare { // ERROR "comparing with != will fail on wrapped errors. Use errors.Is to check for a specific error"
		log.Println("not errCompare")
	}
	if errCompare == err { // ERROR "comparing with == will fail on wrapped errors. Use errors.Is to check for a specific error"
		log.Println("errCompare")
	}
	if errCompare != err { // ERROR "comparing with != will fail on wrapped errors. Use errors.Is to check for a specific error"
		log.Println("not errCompare")
	}
	switch err { // ERROR "switch on an error will fail on wrapped errors. Use errors.Is to check for specific errors"
	case errCompare:
		log.Println("errCompare")
	}
	switch errorLintDoThing() { // ERROR "switch on an error will fail on wrapped errors. Use errors.Is to check for specific errors"
	case errCompare:
		log.Println("errCompare")
	}
}
