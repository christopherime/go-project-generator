package generator

import (
	"github.com/christopherime/go-project-generator/internal/helpers"
	"os"
)

func GenerateCommonFile(projectName string) error {

	// run go mod init projectName
	err := helpers.ExecuteCmd("go", []string{"mod", "init", projectName})
	if err != nil {
		panic("Error running go mod init " + err.Error())
	}

	// Create the gitignore file
	f, err := os.Create(projectName + "/.gitignore")
	if err != nil {
		return err
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			panic("Error closing file " + err.Error())
		}
	}(f)

	// Populate the gitignore file
	_, err = f.WriteString(gitignoreTemplate)
	if err != nil {
		panic("Error writing to main.go file " + err.Error())
	}

	return nil
}
