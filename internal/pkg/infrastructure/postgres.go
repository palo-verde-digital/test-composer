package infrastructure

type Postgres struct {
	Enabled bool   `yaml:"enabled"`
	Tag     string `yaml:"tag"`
}
