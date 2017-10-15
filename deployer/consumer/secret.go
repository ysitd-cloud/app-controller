package consumer

import (
	"github.com/streadway/amqp"
	"github.com/ysitd-cloud/app-controller/core"
)

func CreateSecret(deployer *core.KubernetesDeployer, deliveries <-chan amqp.Delivery) {
	for delivery := range deliveries {
		id := string(delivery.Body)
		if err := deployer.CreateSecret(id); err != nil {
			delivery.Reject(true)
		} else {
			delivery.Ack(false)
		}
	}
}

func UpdateSecret(deployer *core.KubernetesDeployer, deliveries <-chan amqp.Delivery) {
	for delivery := range deliveries {
		id := string(delivery.Body)
		if err := deployer.UpdateSecret(id); err != nil {
			delivery.Reject(true)
		} else {
			delivery.Ack(false)
		}
	}
}

func DeleteSecret(deployer *core.KubernetesDeployer, deliveries <-chan amqp.Delivery) {
	for delivery := range deliveries {
		id := string(delivery.Body)
		if err := deployer.DeleteSecret(id); err != nil {
			delivery.Reject(true)
		} else {
			delivery.Ack(false)
		}
	}
}
