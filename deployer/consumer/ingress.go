package consumer

import (
	"github.com/streadway/amqp"
	"github.com/ysitd-cloud/app-controller/core"
)

func CreateIngress(deployer *core.KubernetesDeployer, deliveries <-chan amqp.Delivery) {
	for delivery := range deliveries {
		id := string(delivery.Body)
		if err := deployer.CreateIngress(id); err != nil {
			delivery.Reject(true)
		} else {
			delivery.Ack(false)
		}
	}
}

func UpdateIngress(deployer *core.KubernetesDeployer, deliveries <-chan amqp.Delivery) {
	for delivery := range deliveries {
		id := string(delivery.Body)
		if err := deployer.UpdateIngress(id); err != nil {
			delivery.Reject(true)
		} else {
			delivery.Ack(false)
		}
	}
}

func DeleteIngress(deployer *core.KubernetesDeployer, deliveries <-chan amqp.Delivery) {
	for delivery := range deliveries {
		id := string(delivery.Body)
		if err := deployer.DeleteIngress(id); err != nil {
			delivery.Reject(true)
		} else {
			delivery.Ack(false)
		}
	}
}
