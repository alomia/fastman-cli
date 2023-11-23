/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// execCmd represents the exec command
var execCmd = &cobra.Command{
	Use:   "exec [command]",
	Short: "Execute a custom command defined in fastman.yaml file",
	Long: `The Exec command allows you to execute custom commands defined in the fastman.yaml configuration file.
	For example:
		$ fastman exec run`,
	Run: func(cmd *cobra.Command, args []string) {
		var defaultCommand string
		if len(args) > 0 {
			defaultCommand = args[0]
		}

		commandMap := viper.GetStringMapString("scripts")

		customizedComman, found := commandMap[defaultCommand]

		if !found {
			fmt.Printf("Error: The command '%s' is not found in the script configuration.\n", defaultCommand)
			return
		}

		commandParts := strings.Fields(customizedComman)

		executable := commandParts[0]
		arguments := commandParts[1:]

		command := exec.Command(executable, arguments...)
		command.Stdout = os.Stdout

		if err := command.Run(); err != nil {
			fmt.Println("could not run command: ", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(execCmd)
}
