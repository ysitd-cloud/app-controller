package consumer

import (
	"github.com/streadway/amqp"
	"github.com/tonyhhyip/go-di-container"
	"github.com/ysitd-cloud/app-controller/core"
)

func StartRoutine(
	app container.Kernel,
	ch *amqp.Channel,
	name string,
	routine func(deployer *core.KubernetesDeployer, deliveries <-chan amqp.Delivery),
) {
	queue, err := ch.QueueDeclare(
		name,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		panic(err)
	}

	err = ch.Qos(1, 0, false)
	if err != nil {
		panic(err)
	}

	messages, err := ch.Consume(
		queue.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		panic(err)
	}
	deployer := app.Make("core.deployer").(*core.KubernetesDeployer)
	go routine(deployer, messages)
}
