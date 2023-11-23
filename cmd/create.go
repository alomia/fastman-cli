package cmd

import (
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create [command]",
	Short: "Create a new project or perform related actions",
	Long: `Create is a command for generating a new project or performing related actions.

For example, you can use the following command to create a new project:
    fastman create project --fastapi`,
}

func init() {
	rootCmd.AddCommand(createCmd)
}
