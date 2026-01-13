package project

import (
	"errors"
	"io"
	"mime/multipart"
	"os"
	"strings"

	"github.com/palo-verde-digital/test-composer/internal/pkg/infrastructure"
	"github.com/palo-verde-digital/test-composer/internal/pkg/testcase"
	"gopkg.in/yaml.v3"
)

var (
	OpenProject *Project
)

type Project struct {
	Name           string                         `yaml:"name"`
	Apps           map[string]Application         `yaml:"apps"`
	Tests          map[string]testcase.TestCase   `yaml:"tests"`
	Infrastructure *infrastructure.Infrastructure `yaml:"infrastructure"`
}

func Create(projectName string) {

	OpenProject = &Project{
		Name:  projectName,
		Apps:  make(map[string]Application),
		Tests: make(map[string]testcase.TestCase),
		Infrastructure: &infrastructure.Infrastructure{
			Postgres: &infrastructure.Postgres{
				Enabled: false,
			},
			Kafka: &infrastructure.Kafka{
				Enabled: false,
				Topics:  make(map[string]string),
			},
			Redis: &infrastructure.Redis{
				Enabled: false,
			},
		},
	}

}

func Read(fileInfo *multipart.FileHeader) error {

	filename := fileInfo.Filename
	if !strings.HasSuffix(filename, ".yaml") && !strings.HasSuffix(filename, ".yml") {
		return errors.New("project file must be .yml/.yaml")
	}

	projectFile, err := fileInfo.Open()
	if err != nil {
		return errors.New("unable to open project file")
	}

	fileContent, err := io.ReadAll(projectFile)
	if err != nil {
		return errors.New("unable to read project file")
	}

	p := &Project{}
	err = yaml.Unmarshal(fileContent, p)
	if err != nil {
		return errors.New("unable to read project data")
	}

	OpenProject = p
	projectFile.Close()

	return nil

}

func Write() error {

	filename := strings.ToLower(OpenProject.Name)
	filename = strings.ReplaceAll(filename, " ", "_") + ".yaml"

	dir, err := os.Getwd()
	if err != nil {
		return errors.New("unable to get working directory")
	}

	filePath := strings.ReplaceAll(dir, "\\", "/") + "/" + filename

	fileContent, err := yaml.Marshal(OpenProject)
	if err != nil {
		return errors.New("unable to marshal to yaml")
	}

	err = os.WriteFile(filePath, fileContent, 0644)
	if err != nil {
		return errors.New("unable to write project file")
	}

	OpenProject = nil
	return nil

}
