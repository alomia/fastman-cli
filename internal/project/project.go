package project

import (
	"path/filepath"

	"github.com/alomia/fastman-cli/internal/fileutil"
	"github.com/alomia/fastman-cli/internal/project/template"
)

// ProjectStructure represents the structure of a project with directories and files.
type ProjectStructure interface {
	Create(basePath string) error
}

// projectStructure defines the internal structure of a project with directories and files.
type projectStructure map[string][]string

// Project represents a project with a specific template and structure.
type project struct {
	template  template.TemplateMap
	structure projectStructure
}

// Create creates the project structure in the base directory.
func (p project) Create(basePath string) error {
	for directory, files := range p.structure {
		if err := fileutil.CreateDirectory(basePath, directory); err != nil {
			return err
		}

		for _, file := range files {
			fullPath := filepath.Join(basePath, directory)
			content := p.template.Get(file)
			if err := fileutil.CreateFile(fullPath, file, content); err != nil {
				return err
			}
		}
	}

	return nil
}
