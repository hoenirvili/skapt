package Skapt

//OptionNotFound
func OptionNotFound(a *App) {
	for i := 0; i < len(a.options); i++ {
		for j := 0; j < len(a.args); j++ {
			if a.options[i].name != a.args[i] {
				panic("Command not found!")
			}
		}
	}
}
