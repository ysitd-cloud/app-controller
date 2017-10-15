package consumer

import (
	"github.com/streadway/amqp"
	"github.com/ysitd-cloud/app-controller/core"
)

func CreateService(deployer *core.KubernetesDeployer, deliveries <-chan amqp.Delivery) {
	for delivery := range deliveries {
		id := string(delivery.Body)
		if err := deployer.CreateService(id); err != nil {
			delivery.Reject(true)
		} else {
			delivery.Ack(false)
		}
	}
}

func UpdateService(deployer *core.KubernetesDeployer, deliveries <-chan amqp.Delivery) {
	for delivery := range deliveries {
		id := string(delivery.Body)
		if err := deployer.UpdateService(id); err != nil {
			delivery.Reject(true)
		} else {
			delivery.Ack(false)
		}
	}
}

func DeleteService(deployer *core.KubernetesDeployer, deliveries <-chan amqp.Delivery) {
	for delivery := range deliveries {
		id := string(delivery.Body)
		if err := deployer.DeleteService(id); err != nil {
			delivery.Reject(true)
		} else {
			delivery.Ack(false)
		}
	}
}
