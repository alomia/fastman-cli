package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/alomia/fastman-cli/internal/fileutil"
	"github.com/alomia/fastman-cli/internal/projectstructure"
	"github.com/spf13/cobra"
)

var fastapiVar bool

// projectCmd represents the project command
var projectCmd = &cobra.Command{
	Use:   "project",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		targetFolder := "."
		if len(args) > 0 {
			targetFolder = args[0]
		}

		fullPathDir := filepath.Join(app.currentDir, targetFolder)

		if fastapiVar {
			if !fileutil.DirectoryExists(app.currentDir, targetFolder) {
				if err := fileutil.CreateDirectory(app.currentDir, targetFolder); err != nil {
					fmt.Printf("Error creating the target directory: %v\n", err)
					return
				}
			}

			fastapi := projectstructure.FastAPI{
				ConfigFile: app.configFile,
			}

			if err := fastapi.Create(fullPathDir); err != nil {
				fmt.Printf("Error when creating FastAPI project %v\n", err)
			}

			fmt.Println("Project structure created successfully.")
		}
	},
}

func init() {
	createCmd.AddCommand(projectCmd)
	projectCmd.Flags().BoolVar(&fastapiVar, "fastapi", false, "create a fastapi project")
}
