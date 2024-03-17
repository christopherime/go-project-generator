package generator

import "os"

func NewPrometheusProject(projectName string) error {
	if err := os.Mkdir(projectName, 0755); err != nil {
		return err
	}

	return nil
}
