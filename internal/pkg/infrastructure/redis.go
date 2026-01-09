package infrastructure

type Redis struct {
	Enabled bool   `yaml:"enabled"`
	Tag     string `yaml:"tag"`
}
