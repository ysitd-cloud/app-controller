package deployer

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

var deletePolicy = metav1.DeletePropagationBackground
var deleteOptions = &metav1.DeleteOptions{
	PropagationPolicy: &deletePolicy,
}

type Controller struct {
	client    kubernetes.Interface
	namespace string
}

func (c *Controller) SetClient(client kubernetes.Interface) {
	c.client = client
}

func (c *Controller) SetNamespace(namespace string) {
	c.namespace = namespace
}
