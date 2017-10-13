package kubernetes

import (
	"fmt"

	"k8s.io/api/core/v1"
	"k8s.io/client-go/pkg/apis/apps/v1beta1"
)

func int32Ptr(i int32) *int32 { return &i }

func (d *deploymentV1) getEnv() []v1.EnvVar {
	vars := make([]v1.EnvVar, 0)

	selector := v1.LocalObjectReference{
		Name: d.GetName(),
	}

	for k, _ := range d.application.GetEnvironment() {
		source := v1.SecretKeySelector{
			LocalObjectReference: selector,
			Key:                  k,
		}

		env := v1.EnvVar{
			Name: k,
			ValueFrom: &v1.EnvVarSource{
				SecretKeyRef: &source,
			},
		}

		vars = append(vars, env)
	}

	return vars
}

func (d *deploymentV1) GetContainer() v1.Container {
	meta := d.application.GetMeta()

	image := fmt.Sprintf("%s:%s", meta.GetImage(), meta.GetTag())

	port := v1.ContainerPort{
		Name:          "http",
		ContainerPort: d.application.GetNetwork().GetPort(),
		Protocol:      v1.ProtocolTCP,
	}

	container := v1.Container{
		Name:  d.GetName(),
		Image: image,
		Ports: []v1.ContainerPort{port},
		Env:   d.getEnv(),
	}

	return container
}

func (d *deploymentV1) GetDeployment() *v1beta1.Deployment {

	container := d.GetContainer()

	template := v1.PodTemplateSpec{
		ObjectMeta: d.getObjectMeta(),
		Spec: v1.PodSpec{
			Containers: []v1.Container{container},
		},
	}

	spec := v1beta1.DeploymentSpec{
		Replicas: int32Ptr(d.application.GetAutoScale().GetReplicas()),
		Template: template,
	}

	return &v1beta1.Deployment{
		ObjectMeta: d.getObjectMeta(),
		Spec:       spec,
	}
}
