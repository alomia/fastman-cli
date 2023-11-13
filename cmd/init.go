package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/alomia/fastman-cli/internal/fileutil"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init [directory]",
	Short: "Create a configuration file",
	Long: `Initialize the project by creating a configuration file in the specified directory.
If [directory] is not provided, the current directory will be used.
	
Examples:
	fastman init
	fastman init directory-name`,
	Run: func(cmd *cobra.Command, args []string) {
		targetFolder := "."
		if len(args) > 0 {
			targetFolder = args[0]
		}

		configFilePath := filepath.Join(app.currentDir, targetFolder)

		if !fileutil.DirectoryExists(app.currentDir, targetFolder) {
			if err := fileutil.CreateDirectory(app.currentDir, targetFolder); err != nil {
				fmt.Printf("Error creating the target directory: %v\n", err)
				return
			}
		}

		err := fileutil.CreateFile(configFilePath, app.configFile)
		if err != nil {
			fmt.Printf("Error creating the configuration file: %v\n", err)
			return
		}

		fmt.Println("Configuration file created at:", configFilePath)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
