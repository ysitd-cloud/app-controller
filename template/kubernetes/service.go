package kubernetes

import (
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/pkg/api/v1"
)

func getPortData(port int32) intstr.IntOrString {
	return intstr.IntOrString{
		Type:   intstr.Int,
		IntVal: port,
	}
}

func (d *deploymentV1) GetService() *v1.Service {
	selector := d.getLabels()

	ports := []v1.ServicePort{
		{
			Name:       "http",
			Port:       80,
			TargetPort: getPortData(d.application.GetNetwork().GetPort()),
			Protocol:   v1.ProtocolTCP,
		},
	}

	spec := v1.ServiceSpec{
		Type:     v1.ServiceTypeClusterIP,
		Selector: selector,
		Ports:    ports,
	}

	return &v1.Service{
		ObjectMeta: d.getObjectMeta(),
		Spec:       spec,
	}
}
