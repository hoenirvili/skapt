package Skapt

import (
	"fmt"
	"testing"
)

func TestApplicatin(t *testing.T) {

	app := New()
	app.SetName("Golang\n")
	app.SetUsage("Usage: \n This is just an echo program simple app\n For displaying content and such things\n")
	app.SetDescription("Description dasjigsafsahgosafushaf")
	app.SetVersion(true, "")
	// Name of the App
	fmt.Println("Name: " + app.GetName())
	// Usage of the program
	fmt.Println(app.GetUsage())
	// Description of te program
	fmt.Println(app.GetDescription())
	// Get the description of the program
	fmt.Println("Version: " + app.GetVersion())

	// oS args
	fmt.Printf("Args off applicaiton provided: %s", app.GetArgs())
	fmt.Println()
}
