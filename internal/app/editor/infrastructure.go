package editor

import (
	"log"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"github.com/palo-verde-digital/test-composer/internal/app/window"
	"github.com/palo-verde-digital/test-composer/internal/pkg/project"
)

func readInfrastucture(c echo.Context) error {

	log.Print("start - infrastructure.readInfrastructure")
	defer log.Print("end - infrastructure.readInfrastructure")

	return c.Render(200, window.InfraTemplateName, nil)

}

func readPostgres(c echo.Context) error {

	log.Print("start - infrastructure.readPostgres")
	defer log.Print("end - infrastructure.readPostgres")

	return c.Render(200, window.PostgresTemplateName, project.OpenProject.Infrastructure.Postgres)

}

func updatePostgres(c echo.Context) error {

	log.Print("start - infrastructure.updatePostgres")
	defer log.Print("end - infrastructure.updatePostgres")

	postgresEnabled := c.FormValue("postgres-enabled") == "on"
	project.OpenProject.Infrastructure.Postgres.Enabled = postgresEnabled

	postgresTag := c.FormValue("postgres-tag")
	project.OpenProject.Infrastructure.Postgres.Tag = postgresTag

	return readPostgres(c)

}

func readKafka(c echo.Context) error {

	log.Print("start - infrastructure.readKafka")
	defer log.Print("end - infrastructure.readKafka")

	return c.Render(200, window.KafkaTemplateName, project.OpenProject.Infrastructure.Kafka)

}

func updateKafka(c echo.Context) error {

	log.Print("start - infrastructure.updateKafka")
	defer log.Print("end - infrastructure.updateKafka")

	kafkaEnabled := c.FormValue("kafka-enabled") == "on"
	project.OpenProject.Infrastructure.Kafka.Enabled = kafkaEnabled

	kafkaTag := c.FormValue("kafka-tag")
	project.OpenProject.Infrastructure.Kafka.Tag = kafkaTag

	kafkaTopic := c.FormValue("kafka-topic")
	if kafkaTopic != "" {
		project.OpenProject.Infrastructure.Kafka.Topics[uuid.NewString()] = kafkaTopic
	}

	return readKafka(c)

}

func updateKafkaTopic(c echo.Context) error {

	log.Print("start - infrastructure.updateKafkaTopic")
	defer log.Print("end - infrastructure.updateKafkaTopic")

	topicId := c.Param("topicId")
	topicName := c.FormValue("kafka-topic-" + topicId)

	project.OpenProject.Infrastructure.Kafka.Topics[topicId] = topicName

	return readKafka(c)

}

func deleteKafkaTopic(c echo.Context) error {

	log.Print("start - infrastructure.deleteKafkaTopic")
	defer log.Print("end - infrastructure.deleteKafkaTopic")

	topicId := c.Param("topicId")
	delete(project.OpenProject.Infrastructure.Kafka.Topics, topicId)

	return readKafka(c)

}

func readRedis(c echo.Context) error {

	log.Print("start - infrastructure.readRedis")
	defer log.Print("end - infrastructure.readRedis")

	return c.Render(200, window.RedisTemplateName, project.OpenProject.Infrastructure.Redis)

}

func updateRedis(c echo.Context) error {

	log.Print("start - infrastructure.updateKafka")
	defer log.Print("end - infrastructure.updateKafka")

	redisEnabled := c.FormValue("redis-enabled") == "on"
	project.OpenProject.Infrastructure.Redis.Enabled = redisEnabled

	redisTag := c.FormValue("redis-tag")
	project.OpenProject.Infrastructure.Redis.Tag = redisTag

	return readRedis(c)

}
