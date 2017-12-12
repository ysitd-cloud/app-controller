package deployer

import "k8s.io/apimachinery/pkg/apis/meta/v1"

var deletePolicy = v1.DeletePropagationBackground
var deleteOptions = &v1.DeleteOptions{
	PropagationPolicy: &deletePolicy,
}
