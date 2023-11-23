package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const version = "0.1.0"
const cfgFile = "fastman.yaml"

var currentDir string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "fastman",
	Version: version,
	Short:   "FastMan CLI: Streamlining Project Initiation and Automation",
	Long: `FastMan is a command-line tool designed to streamline project initiation and automation.

Key Features:
	- Rapid Project Structure Generation: Efficiently create a project's structure, eliminating the need for manual configurations, repetitive directories, and lengthy commands.
	- Command Automation with Aliases: Expedite the execution of common tasks by automating frequent commands through customizable aliases.
	
Why FastMan?
	FastMan arises from the need to expedite the process of project initiation and management.
Initial setup and repetitive command execution can consume time and energy.
FastMan is crafted to eliminate these barriers, allowing you to focus on what truly matters: your project.
	
Usage Examples:
	$ fastman create project --fastapi
	
Discover how FastMan can enhance your workflow and make project management more efficient.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	setCurrentDir()
}

func initConfig() {
	viper.AddConfigPath(currentDir)
	viper.SetConfigType("yaml")
	viper.SetConfigName("fastman")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error config file: The file %s does not exist in the directory\n", cfgFile)
		return
	}
}

func setCurrentDir() {
	var err error

	currentDir, err = os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
}
