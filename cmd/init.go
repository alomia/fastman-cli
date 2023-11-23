package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/alomia/fastman-cli/internal/fileutil"
	"github.com/alomia/fastman-cli/internal/project/template"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create a configuration file",
	Long: `Initializes the project by creating a configuration file in the current directory.
	
Examples:
	fastman init`,
	Run: func(cmd *cobra.Command, args []string) {
		fastmanTemplate := template.Fastman.Get(cfgFile)

		err := fileutil.CreateFile(currentDir, cfgFile, fastmanTemplate)
		if err != nil {
			fmt.Printf("Error creating the configuration file: %v\n", err)
			return
		}

		fmt.Printf("Configuration file created at: %s\n", filepath.Join(currentDir, cfgFile))
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
