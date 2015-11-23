package Skapt

import (
	"fmt"
	"testing"
)

func TestApplicatin(t *testing.T) {
	// ====================== INIT ===========================
	var f = [][]string{
		[]string{
			"-f",
			"--full-path",
			"-g",
			"-c",
			"-hf",
		},
		[]string{
			"-l",
			"--link-pink",
			"-f",
			"-k",
		},
		[]string{
			"-d",
			"--direct-pink",
		}}

	var h = []Handler{
		func() {
			fmt.Println("func 1")
		},
		func() {
			fmt.Println("func 2")
		},
		func() {
			fmt.Println("func 3")
		}}

	app := NewApp()

	app.SetName("Golang\n")

	app.SetUsage("Usage: \n This is just an echo program simple app\n For displaying content and such things\n")
	app.SetDescription("Description dasjigsafsahgosafushaf")
	app.SetVersion(true, "")
	// Set OPTIONS
	app.SetOptions(f, h)

	// ====================== PRINT ==========================

	// Name of the App
	fmt.Println("Name: " + app.Name())

	// Usage of the program
	fmt.Println(app.Usage())

	// Description of te program
	fmt.Println(app.Description())

	// Options
	fmt.Println(app.Options())
	for _, val := range app.options {
		val.Run()
	}

	// Get the description of the program
	fmt.Println("Version: " + app.Version())
	// set the command for the application
	fmt.Printf("Args off applicaiton provided: %s", app.Args())
	fmt.Println()
}
