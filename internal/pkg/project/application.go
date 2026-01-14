package project

import (
	"errors"
	"os"
	"path"
	"strings"
	"unicode"

	"github.com/palo-verde-digital/test-composer/pkg/stringutil"
)

type Application struct {
	Name    string              `yaml:"name"`
	Image   string              `yaml:"image"`
	IsApi   bool                `yaml:"isApi"`
	ApiPort string              `yaml:"apiPort"`
	Env     map[string]Variable `yaml:"env"`
}

type Variable struct {
	Key   string `yaml:"key"`
	Value string `yaml:"val"`
}

func CreateApplication() Application {

	return Application{
		Env: make(map[string]Variable),
	}

}

func ValidateApplicationName(appName string) string {

	if appName == "" || strings.TrimSpace(appName) == "" {
		return "cannot be blank"
	}

	return ""

}

func ValidateApplicationImage(appId, appImage string) string {

	if appImage == "" || strings.TrimSpace(appImage) == "" {
		return "cannot be blank"
	} else if strings.ToUpper(appImage) == "DOCKERFILE" {
		for id, app := range OpenProject.Apps {
			if app.Image == "DOCKERFILE" && id != appId {
				return "duplicate usage of local Dockerfile"
			}
		}

		if wd, err := os.Getwd(); err != nil {
			return "unknown error occurred"
		} else if _, err := os.Stat(path.Join(wd, "Dockerfile")); err != nil {
			if errors.Is(err, os.ErrNotExist) {
				return "no local Dockerfile"
			} else {
				return "unknown error occurred"
			}
		}
	} else if !strings.Contains(appImage, ":") {
		return "must be 'DOCKERFILE', or of format IMAGE_NAME:TAG"
	} else if !unicode.IsLetter(rune(appImage[0])) {
		return "must start w/ a lowercase letter"
	} else if stringutil.ContainsUpper(appImage) {
		return "cannot contain uppercase letters"
	}

	return ""

}

func ValidateApplicationApiPort(appApiPort string) string {

	if !stringutil.IsNumeric(appApiPort) {
		return "must be numeric value"
	}

	return ""

}
