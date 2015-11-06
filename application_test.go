package Skapt

import (
	"fmt"
	"testing"
)

func TestApplicatin(t *testing.T) {
	var myOptions = []string{"g", "h", "i"}
	var handlers []FlagFunc

	firstHandler := func() {
		fmt.Println("FirstMethod Test")
	}
	secondHandler := func() {
		fmt.Println("Another bullshit test")
	}
	thirdHandler := func() {
		fmt.Println("The last bullshit test")
	}
	handlers = append(handlers, firstHandler, secondHandler, thirdHandler)
	app := New()
	app.SetName("Golang\n")
	app.SetUsage("Usage: \n This is just an echo program simple app\n For displaying content and such things\n")
	app.SetDescription("Description dasjigsafsahgosafushaf")
	app.SetVersion(true, "")
	app.SetNameOptions(myOptions)
	app.SetOptionHandlers(handlers)
	// Name of the App
	fmt.Println("Name: " + app.GetName())
	// Usage of the program
	fmt.Println(app.GetUsage())
	// Description of te program
	fmt.Println(app.GetDescription())
	// Get the description of the program
	fmt.Println("Version: " + app.GetVersion())
	// Get all the flags
	std := app.GetNameOptions()
	fmt.Printf("Options: ")
	if std != nil {
		for _, val := range std {
			fmt.Printf("%s ", val)
		}
	}
	fmt.Println("Test handler for every flag")
	for _, val := range app.options {
		fmt.Printf("%s ", val.name)
		val.Run()
	}
	fmt.Println()
}
