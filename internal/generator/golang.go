package generator

import "os"

func NewGolangProject(projectName string) error {
	if err := os.Mkdir(projectName, 0755); err != nil {
		return err
	}

	return nil
}
