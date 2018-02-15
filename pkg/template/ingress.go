package template

import "k8s.io/api/extensions/v1beta1"

func GenerateIngress(id, domain string) *v1beta1.Ingress {
	backend := v1beta1.IngressBackend{
		ServiceName: GetName(id),
		ServicePort: getPortData(80),
	}

	rulePath := v1beta1.HTTPIngressPath{
		Path:    "/",
		Backend: backend,
	}

	ruleValue := v1beta1.HTTPIngressRuleValue{
		Paths: []v1beta1.HTTPIngressPath{rulePath},
	}

	rule := v1beta1.IngressRule{
		Host: domain,
		IngressRuleValue: v1beta1.IngressRuleValue{
			HTTP: &ruleValue,
		},
	}

	spec := v1beta1.IngressSpec{
		Rules: []v1beta1.IngressRule{rule},
	}

	return &v1beta1.Ingress{
		ObjectMeta: getObjectMeta(id),
		Spec:       spec,
	}
}
