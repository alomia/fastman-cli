package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "fastman",
	Short: "FastMan CLI: Streamlining Project Initiation and Automation",
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
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
