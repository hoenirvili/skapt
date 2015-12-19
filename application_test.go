package Skapt

import (
	"fmt"
	"os"
	"testing"
)

// ====================== INIT ===========================
/*var f = [][]string{
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
**/
var h = []Handler{
	func() {
		fmt.Println("func 0")
	},
	func() {
		fmt.Println("func 1")
	},
	func() {
		fmt.Println("func 2")
	},
	func() {
		fmt.Println("func 3")
	}}

//========================================================

func TestApplicatin(t *testing.T) {
	os.Args = []string{"", "-k", "-C", "-f", "--path=file/to/file/peer"}
	app := NewApp()
	app.SetName("Golang\n")

	app.SetUsage("Usage: \n This is just an echo program simple app\n For displaying content and such things\n")
	app.SetDescription("Description dasjigsafsahgosafushaf")
	app.SetVersion(true, "")
	// Set OPTIONS

	// Declare new commands
	//app.AppendNewCommand("test1", "test2", "test3", f, h)
	//app.AppendNewCommand("q1", "q2", "q3", f, h)
	//app.AppendNewCommand("r1", "r2", "r3", f, h)
	app.AppenNewOption("-f", "-FUL", nil, h[0])
	app.AppenNewOption("-C", "--create", []string{"--path="}, h[1])
	app.AppenNewOption("-G", "--mik", nil, h[2])
	app.AppenNewOption("-k", "", nil, h[3])
	// Declare new options

	// ====================== RUN ============================
	fmt.Println("================")
	app.Run()
	fmt.Println("================")
	// ====================== PRINT ==========================

	// Name of the App
	fmt.Println("Name: " + app.Name())

	// Usage of the program
	fmt.Println(app.Usage())

	// Description of te program
	fmt.Println(app.Description())

	// Options
	if app.Options != nil {
		for _, val := range app.options {
			fmt.Println(val.Name())
			fmt.Println(val.Alias())
			fmt.Println(val.RequireFlags())
		}
	}

	// Print Commands
	if app.commands != nil {
		for _, cmd := range app.commands {
			fmt.Println(cmd.Name())
			fmt.Println(cmd.Usage())
			fmt.Println(cmd.Description())
			for _, opt := range cmd.options {
				fmt.Println(opt)
			}
		}
	}

	// Get the description of the program
	fmt.Println("Version: " + app.Version())
	// set the command for the application
	fmt.Printf("Args off applicaiton provided: %s", app.Args())
	fmt.Println()
}
