package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	// "os/exec"
	"strings"
)

var fullPath string

func usePath (path string) (string , error) {
		if strings.Contains(path, "\\") {
			return path , nil
		} else {
			initialPath, errPath := os.Getwd()

			if errPath != nil {
				return "" , errors.New("Error getting your current path")
			}

			fullPath = initialPath + "\\" + path

			return path , nil

		}

}

func main() {

	// Read path from command line

	if len(os.Args) < 2 {
		fmt.Println("Please provide a path")
		os.Exit(1)
	} else if len(os.Args) > 2 {
		fmt.Println("Too many arguments")
		os.Exit(1)
	} else {

	path , errorPath := usePath(os.Args[1])

	if errorPath != nil {
		fmt.Println(errorPath)
		os.Exit(1)
	}

	file, errFile := os.Open(path)

		if errFile != nil {
			fmt.Println(errFile)
		}

		defer file.Close()

		//  Change the file from bits to text

		text, errText := ioutil.ReadAll(file)

		if errText != nil {
			fmt.Println(errText)
		}

		fmt.Print(string(text))

		// exec.Command("cmd", "/c", "start", path).Run()

		os.Exit(0)

	}

		
}

