package infrastructure

type Infrastructure struct {
	Postgres *Postgres `yaml:"postgres"`
	Kafka    *Kafka    `yaml:"kafka"`
	Redis    *Redis    `yaml:"redis"`
}
