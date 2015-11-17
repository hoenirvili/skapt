package Skapt

import (
	"fmt"
	"testing"
)

func TestApplicatin(t *testing.T) {

	app := New(false)
	app.SetName("Golang\n")
	app.SetUsage("Usage: \n This is just an echo program simple app\n For displaying content and such things\n")
	app.SetDescription("Description dasjigsafsahgosafushaf")
	app.SetVersion(true, "")
	// Name of the App
	fmt.Println("Name: " + app.Name())
	// Usage of the program
	fmt.Println(app.Usage())
	// Description of te program
	fmt.Println(app.Description())
	// Get the description of the program
	fmt.Println("Version: " + app.Version())
	// Get the Mode of the app
	if app.Mode() {
		fmt.Println("App Mode : true")
	} else {
		fmt.Println("App Mode : false")
	}
	// set the command for the application
	containerCommands := make([][]string, 3)
	for i := range containerCommands {
		containerCommands[i] = make([]string, 3)
	}
	// oS args
	fmt.Printf("Args off applicaiton provided: %s", app.Args())
	fmt.Println()
}
