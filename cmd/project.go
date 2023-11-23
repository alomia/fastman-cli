package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/alomia/fastman-cli/internal/fileutil"
	"github.com/alomia/fastman-cli/internal/project"
	"github.com/spf13/cobra"
)

var fastapiVar bool
var nameVar string

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
		if fastapiVar {
			fastapi := project.FastAPI{}

			if len(nameVar) > 0 {
				err := fileutil.CreateDirectory(currentDir, nameVar)
				if err != nil {
					fmt.Printf("Error creating %s directory: %v", nameVar, err)
					return
				}

				currentDir = filepath.Join(currentDir, nameVar)
			}

			if err := fastapi.Create(currentDir); err != nil {
				fmt.Printf("Error creating FastAPI project structure: %v\n", err)
				return
			}

			fmt.Println("Project structure created successfully")
		}
	},
}

func init() {
	projectCmd.Flags().BoolVar(&fastapiVar, "fastapi", false, "create a fastapi project")
	projectCmd.Flags().StringVarP(&nameVar, "name", "n", "", "project name")

	createCmd.AddCommand(projectCmd)
}
