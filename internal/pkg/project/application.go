package project

type Application struct {
	Image string              `yaml:"image"`
	Env   map[string]Variable `yaml:"env"`
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
