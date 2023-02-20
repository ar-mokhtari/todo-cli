package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	separator = strings.Repeat("-", 20)
	Scanner   = bufio.NewScanner(os.Stdin)
)

func ShowLogo() (logo string) {
	return "     _    _ _       _____     ____        \n    / \\  | (_)____ |_   _|__ |  _ \\  ___  \n   / _ \\ | | |_  /   | |/ _ \\| | | |/ _ \\ \n  / ___ \\| | |/ /    | | (_) | |_| | (_) |\n /_/   \\_\\_|_/___|   |_|\\___/|____/ \\___/ \n                                          "
}

func ShowMenu() (output string) {
	output += fmt.Sprintln("***\tMain menu:\t***\n", "Please select:")
	output += fmt.Sprintln(separator)
	output += fmt.Sprintln("***\t\t CATEGORY \t\t***")
	output += fmt.Sprintln("***\t:\t***")
	return output
}

func RunMenu() {
	fmt.Println(ShowLogo())
	for {
		if UserAuth == nil {
			if err := SignIn(); err != nil {
				continue
			}
		} else {
			fmt.Println(ShowMenu())
			for {
				Scanner.Scan()
				switch Scanner.Text() {
				case "category":
					println("ssdasdasd")
					fmt.Println(ShowMenu())
				default:
					fmt.Println("please select another command")
					fmt.Println(ShowMenu())
				}
			}
		}
	}
}
