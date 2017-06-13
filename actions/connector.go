package actions

import (
	"fmt"
	"log"

	"github.com/vice-registry/vice-api/models"

	"github.com/streadway/amqp"
)

var rabbitmqCredentials = struct {
	Location string
	Username string
	Password string
	Channel  *amqp.Channel
}{}

// SetRabbitmqCredentials set the login credentials to Couchbase cluster
func SetRabbitmqCredentials(location string, username string, password string) {
	rabbitmqCredentials.Location = location
	rabbitmqCredentials.Username = username
	rabbitmqCredentials.Password = password

	conn, err := amqp.Dial("amqp://" + username + ":" + password + "@" + location + "/")
	failOnError(err, "Failed to connect to RabbitMQ")
	//defer conn.Close()

	channel, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	//defer ch.Close()

	rabbitmqCredentials.Channel = channel
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func sendMessage(queueName string, message string) {
	importQueue, err := rabbitmqCredentials.Channel.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	failOnError(err, "Failed to declare a queue "+queueName)

	body := message
	err = rabbitmqCredentials.Channel.Publish(
		"",               // exchange
		importQueue.Name, // routing key
		false,            // mandatory
		false,            // immediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(body),
		})
	failOnError(err, "Failed to publish a message to queue "+queueName)
}

func NewImportAction(image *models.Image) {
	sendMessage("import", image.ID)
}

func NewExportAction(deployment *models.Deployment) {
	sendMessage("export", deployment.ID)
}
