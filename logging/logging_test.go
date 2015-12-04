package logging_test

import (
	"."
)

func Example() {
	l, _ := logging.NewLogger("PROSPERO. ", 0, logging.Stdout)
	l.Print("What see’st thou else / In the dark backward and abysm of time?")
	// Output: PROSPERO. What see’st thou else / In the dark backward and abysm of time?
}
