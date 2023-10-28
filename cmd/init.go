package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/alomia/fastman-cli/internal/fileutil"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "init [targetFolder]",
	Long: `Initialize the project by creating a configuration file in the specified directory.
If [targetFolder] is not provided, the current directory will be used.`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		targetFolder := "."
		if len(args) > 0 {
			targetFolder = args[0]
		}

		currentDir, err := os.Getwd()
		if err != nil {
			fmt.Printf("Error getting the current directory: %v\n", err)
			return
		}

		if !fileutil.DirectoryExists(currentDir, targetFolder) {
			err = fileutil.CreateDirectory(currentDir, targetFolder)
			if err != nil {
				fmt.Printf("Error creating the target directory: %v\n", err)
				return
			}
		}

		configFileName := "fastman.yaml"
		configFilePath := filepath.Join(currentDir, targetFolder, configFileName)

		err = fileutil.CreateFile(configFilePath)
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
