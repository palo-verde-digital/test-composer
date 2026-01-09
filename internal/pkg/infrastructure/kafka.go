package infrastructure

type Kafka struct {
	Enabled bool     `yaml:"enabled"`
	Tag     string   `yaml:"tag"`
	Topics  []string `yaml:"topics"`
}
