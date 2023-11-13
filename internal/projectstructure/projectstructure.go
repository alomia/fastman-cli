package projectstructure

import (
	"fmt"
	"path/filepath"

	"github.com/alomia/fastman-cli/internal/fileutil"
	"github.com/alomia/fastman-cli/internal/projectstructure/template"
)

// ProjectStructureInterface es una interfaz que representa la estructura de un proyecto con directorios y archivos.
type ProjectStructureInterface interface {
	Create(basePath string) error
}

// Structure represents the structure of a project with directories and files.
type ProjectStructure map[string][]string

func (p ProjectStructure) Create(basePath string) error {
	for directory, files := range p {
		if err := fileutil.CreateDirectory(basePath, directory); err != nil {
			return err
		}

		for _, file := range files {
			fullPath := filepath.Join(basePath, directory)
			if err := fileutil.CreateFile(fullPath, file, template.FastAPITemplates[file]); err != nil {
				return err
			}
		}
	}
	return nil
}

type FastAPI struct {
	ConfigFile string
}

func (f *FastAPI) Create(basePath string) error {
	fastapi := ProjectStructure{
		".":        {"main.py", "requirements.txt", f.ConfigFile},
		"database": {"__init__.py", "connection.py"},
		"routes":   {"__init__.py", "events.py", "users.py"},
		"models":   {"__init__.py", "events.py", "users.py"},
	}

	if err := fastapi.Create(basePath); err != nil {
		return fmt.Errorf("Error creating project structure: %w", err)
	}

	return nil
}
