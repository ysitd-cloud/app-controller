package consumer

import (
	"github.com/streadway/amqp"
	"github.com/ysitd-cloud/app-controller/core"
)

func CreateDeplotment(deployer *core.KubernetesDeployer, deliveries <-chan amqp.Delivery) {
	for delivery := range deliveries {
		id := string(delivery.Body)
		if err := deployer.CreateDeployment(id); err != nil {
			delivery.Reject(true)
		} else {
			delivery.Ack(false)
		}
	}
}

func UpdateDeployment(deployer *core.KubernetesDeployer, deliveries <-chan amqp.Delivery) {
	for delivery := range deliveries {
		id := string(delivery.Body)
		if err := deployer.UpdateDeployment(id); err != nil {
			delivery.Reject(true)
		} else {
			delivery.Ack(false)
		}
	}
}

func DeleteDeployment(deployer *core.KubernetesDeployer, deliveries <-chan amqp.Delivery) {
	for delivery := range deliveries {
		id := string(delivery.Body)
		if err := deployer.DeleteDeployment(id); err != nil {
			delivery.Reject(true)
		} else {
			delivery.Ack(false)
		}
	}
}
