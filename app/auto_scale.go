package app

type AutoScale interface {
	GetReplicas() int32
}

func NewAutoScale(replicas int32) AutoScale {
	return &autoScale{
		replicas: replicas,
	}
}

type autoScale struct {
	replicas int32
}

func (a *autoScale) GetReplicas() int32 {
	return a.replicas
}
