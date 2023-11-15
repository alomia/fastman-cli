package project

import "github.com/alomia/fastman-cli/internal/project/template"

// FastAPI represents the specific structure of a FastAPI project.
type FastAPI struct{}

// Create creates the structure of a FastAPI project in the base directory.
func (f FastAPI) Create(basePath string) error {
	fastapi := project{
		template: template.FastAPI,
		structure: projectStructure{
			".":        {"main.py", "requirements.txt"},
			"database": {"__init__.py", "connection.py"},
			"routes":   {"__init__.py", "events.py", "users.py"},
			"models":   {"__init__.py", "events.py", "users.py"},
		},
	}

	if err := fastapi.Create(basePath); err != nil {
		return err
	}

	return nil
}
