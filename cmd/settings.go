package cmd

import (
	"fmt"
	"os"
)

type Application struct {
	version    string
	configFile string
	currentDir string
}

var app *Application

func init() {
	app = &Application{
		version:    "0.1.0",
		configFile: "fastman.yaml",
		currentDir: setCurrentDir(),
	}
}

func setCurrentDir() string {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting the current directory: %v\n", err)
		return ""
	}

	return currentDir
}
