package Skapt

import (
	"fmt"
	"os"
	"testing"
)

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

func TestApplicatin(t *testing.T) {
	os.Args = []string{"", "-k", "-C", "-f", "--path", "file/to/file/peer", "--number", "522"}
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
	app.AppenNewOption("-f", "-FUL", nil, BOOL, h[0])
	app.AppenNewOption("-C", "--create", []string{"--path"}, BOOL, h[1])
	app.AppenNewOption("--path", "-pth", nil, STRING, nil)
	app.AppenNewOption("-G", "--mik", nil, BOOL, h[2])
	app.AppenNewOption("-k", "", nil, BOOL, h[3])
	app.AppenNewOption("--number", "-nr", nil, INT, nil)
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

	// Description of the program
	fmt.Println(app.Description())
	fmt.Println()

	// Options
	if app.Options != nil {
		for _, val := range app.options {
			fmt.Print(val.Name() + " " + val.Alias() + " ")
			fmt.Print(val.RequireFlags())
			if val.TypeFlag() == 0 {
				fmt.Print(" BOOL")
			} else {
				if val.TypeFlag() == 1 {
					fmt.Print(" STRING")
				} else {
					if val.TypeFlag() == 2 {
						fmt.Print(" INT")
					}
				}
			}
			fmt.Println()
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
	fmt.Println()
	fmt.Println()
	fmt.Printf("Args off applicaiton provided: %s", app.Args())
	fmt.Println()
	fmt.Println()
	fmt.Println()
}

//TODO
func TestFlag(t *testing.T) {
	os.Args = []string{"", "-k", "-C", "-f", "-pth", "file/to/file/peer", "--number", "522"}

	app := NewApp()

	app.SetVersion(true, "")

	app.AppenNewOption("-f", "-FUL", nil, BOOL, h[0])
	app.AppenNewOption("-C", "--create", []string{"--path"}, BOOL, h[1])
	app.AppenNewOption("--path", "-pth", nil, STRING, nil)
	app.AppenNewOption("-G", "--mik", nil, BOOL, h[2])
	app.AppenNewOption("-k", "", nil, BOOL, h[3])
	app.AppenNewOption("--number", "-nr", nil, INT, nil)
	// Declare new options

	// ====================== RUN ============================
	fmt.Println("================")
	app.Run()
	fmt.Println("================")
	// ====================== PRINT ==========================

	fmt.Println()
	fmt.Print("--path STRING=")
	fmt.Println(app.String("--path"))
	fmt.Println()
	/*
		fmt.Print("-k BOOL= ")
		fmt.Println(app.Bool("-k"))
		fmt.Print("-G BOOL= ")
		fmt.Println(app.Bool("-G"))
		fmt.Print("-FUL BOOL= ")
		fmt.Println(app.Bool("-FUL"))
		fmt.Print("mmm BOOL= ")
		fmt.Println(app.Bool("mmm"))
		fmt.Print("--number INT= ")
		fmt.Println(app.Int("--number"))
		fmt.Println()
	*/
	// Get the description of the program
	fmt.Println("Version: " + app.Version())
	// set the command for the application
	fmt.Println()
	fmt.Println()
	fmt.Printf("Args off applicaiton provided: %s", app.Args())
	fmt.Println()
	fmt.Println()
	fmt.Println()
}
