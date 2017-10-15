package main

import (
	"os"

	"github.com/streadway/amqp"
	"github.com/ysitd-cloud/app-controller/deployer/consumer"
	"github.com/ysitd-cloud/app-controller/deployer/kernel"
)

func main() {
	url := os.Getenv("AMQP_URL")
	app := kernel.CreateKernel()

	conn, err := amqp.Dial(url)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	forever := make(chan bool)

	consumer.StartRoutine(app, ch, "deployment:create", consumer.CreateDeplotment)
	consumer.StartRoutine(app, ch, "deployment:update", consumer.UpdateDeployment)
	consumer.StartRoutine(app, ch, "deployment:delete", consumer.DeleteDeployment)
	consumer.StartRoutine(app, ch, "ingress:create", consumer.CreateIngress)
	consumer.StartRoutine(app, ch, "ingress:update", consumer.UpdateIngress)
	consumer.StartRoutine(app, ch, "ingress:delete", consumer.DeleteIngress)
	consumer.StartRoutine(app, ch, "service:create", consumer.CreateService)
	consumer.StartRoutine(app, ch, "service:update", consumer.UpdateService)
	consumer.StartRoutine(app, ch, "service:delete", consumer.DeleteService)
	consumer.StartRoutine(app, ch, "secret:create", consumer.CreateSecret)
	consumer.StartRoutine(app, ch, "secret:update", consumer.UpdateSecret)
	consumer.StartRoutine(app, ch, "secret:delete", consumer.DeleteSecret)

	<-forever
}
