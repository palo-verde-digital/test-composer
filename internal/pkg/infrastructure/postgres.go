package infrastructure

type Postgres struct {
	Databases []string `yaml:"databases"`
}
