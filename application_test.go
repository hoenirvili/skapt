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
	},
	func() {
		fmt.Println("func 4")
	},
}

func TestApplication(t *testing.T) {
	os.Args = []string{"", "-k", "-C", "-f", "--path", "file/to/file/peer", "-G"}
	var auth = []string{"Jim Cook", "Alfred Benedict", "G G.", "Hacker Pacer"}
	app := NewApp()
	app.SetName("Golang")

	app.SetUsage("This is just an echo program simple app  For displaying content and such things\n")
	app.SetDescription("Description dasjigsafsahgosafushaf")
	app.SetVersion(true, "")
	app.SetAuthors(auth)
	// Set OPTIONS

	app.AppendNewOption("-f", "-FUL", "LOREM ipsum loremipsjduijdsiauhf iausdhf isadhf isudhf aisuf", BOOL, h[0])
	app.AppendNewOption("-C", "--create", "LOREM ipsum loremipsjduijdsiauhf iausdhf isadhf isudhf aisuf", BOOL, h[1])
	app.AppendNewOption("--path", "-pth", "LOREM ipsum loremipsjduijdsiauhf iausdhf isadhf isudhf aisuf", STRING, nil)
	app.AppendNewOption("-G", "--mik", "LOREM ipsum loremipsjduijdsiauhf iausdhf isadhf isudhf aisuf", BOOL, h[2])
	app.AppendNewOption("-k", "", "LOREM ipsum loremipsjduijdsiauhf iausdhf isadhf isudhf aisuf", BOOL, h[3])
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
	for _, val := range app.options {
		fmt.Print(val.Name() + " " + val.Alias() + " ")
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

func TestFlag(t *testing.T) {
	os.Args = []string{"", "-k", "-C", "-f", "-pth", "file/to/file/peer", "--number", "522"}

	app := NewApp()

	var auth = []string{"Jim Cook", "Alfred Benedict", "G G.", "Hacker Pacer"}

	app.SetUsage("This is just an echo program simple app  For displaying content and such things\n")
	app.SetDescription("Description dasjigsafsahgosafushaf")
	app.SetVersion(true, "")
	app.SetAuthors(auth)

	app.AppendNewOption("-f", "-FUL", "Lorem ipsum joke mmomomdoas osaiudfhsi ughisuadh gas", BOOL, h[0])
	app.AppendNewOption("-C", "--create", "Lorem ipsum ijiojasdofj oasdijfosaidjf", BOOL, h[1])
	app.AppendNewOption("--path", "-pth", "Lorem ipsum ijiojasdofj oasdijfosaidjf", STRING, nil)
	app.AppendNewOption("-G", "--mik", "Lorem ipsum ijiojasdofj oasdijfosaidjf", BOOL, h[2])
	app.AppendNewOption("-k", "", "Lorem ipsum ijiojasdofj oasdijfosaidjf", BOOL, h[3])
	app.AppendNewOption("--number", "-nr", "Lorem ipsum ijiojasdofj oasdijfosaidjf", INT, nil)
	// Declare new options

	// ====================== RUN ============================
	fmt.Println("================")
	app.Run()
	fmt.Println("================")
	// ====================== PRINT ==========================

	fmt.Println()
	fmt.Print("--path STRING=")
	fmt.Println(app.String("--path"))
	fmt.Print("-k BOOL=\t")
	fmt.Println(app.Bool("-k"))
	fmt.Print("-G BOOL=\t")
	fmt.Println(app.Bool("-G"))
	fmt.Print("-FUL BOOL=\t")
	fmt.Println(app.Bool("-FUL"))
	fmt.Print("-mmm BOOL =\t")
	fmt.Println(app.Bool("mmm"))
	fmt.Print("--number INT=\t")
	fmt.Println(app.Int("-nr"))
	fmt.Println()

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
func TestFlagCommand(t *testing.T) {
	os.Args = []string{"", "save", "flag6", "dasd"}
	app := NewApp()
	app.SetName("Golang\n")
	app.SetVersion(false, "1.0.0")
	app.SetAuthors([]string{"Hoenirvili"})
	// Declare new commands
	var f [][]string
	//set all things
	f = append(f, []string{"flag1", "flag2", "Lorem usadiuhdsaiufhsdiuahfisuadhf iusdahf sad", "INT", ""})
	f = append(f, []string{"flag3", "flag4", "Lomrqiudsaihfiashfiasuhf iash fjashf iuahs f", "STRING", ""})
	f = append(f, []string{"flag5", "flag6", "oijadsiusad idsjkldsaloiasdh fhsadf hasd fih dsaifh ", "BOOL", ""})

	app.AppendNewCommand("init", "Init the tests or our application", "Inits all the app", f, h)
	app.AppendNewCommand("save", "Saves all logs.", "Unwanted usage", f, h)
	app.AppendNewCommand("install", "Install the system drivers", "Usage usage", f, h)

	// ====================== RUN ============================
	fmt.Println("================")
	app.Run()
	fmt.Println("================")
	// ====================== PRINT ==========================
	//
	fmt.Println()
	fmt.Println()
	fmt.Printf("Args off applicaiton provided: %s", app.Args())
	fmt.Println()
	fmt.Println()
	fmt.Println()

}
