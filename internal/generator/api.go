package generator

import (
	"github.com/christopherime/go-project-generator/internal/helpers"
	"os"
)

func NewAPIProject(project helpers.Project) error {
	var err error

	// create the path
	err = os.Mkdir(project.Path, 0755)
	if err != nil {
		panic("Error creating project directory " + err.Error())
	}

	return err
}
