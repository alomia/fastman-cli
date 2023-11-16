package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/alomia/fastman-cli/internal/fileutil"
	"github.com/alomia/fastman-cli/internal/project"
	"github.com/spf13/cobra"
)

var fastapiVar bool

// projectCmd represents the project command
var projectCmd = &cobra.Command{
	Use:   "project [arguments] [options]",
	Short: "Manage projects under the 'create' command",
	Long: `Project is a subcommand of 'create' for managing projects with various options and configurations.

For example:
	To create a new project:
		fastman create project --fastapi
		fastman create project planner --fastapi`,
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

			fastapi := project.FastAPI{}

			if err := fastapi.Create(fullPathDir); err != nil {
				fmt.Printf("Error creating FastAPI project structure: %v\n", err)
				return
			}

			fmt.Println("Project structure created successfully.")
		}
	},
}

func init() {
	createCmd.AddCommand(projectCmd)
	projectCmd.Flags().BoolVar(&fastapiVar, "fastapi", false, "create a fastapi project")
}
