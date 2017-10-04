package app

func NewAutoScale(replicas int32) AutoScale {
	return &autoScale{
		replicas: replicas,
	}
}

func (a *autoScale) GetReplicas() int32 {
	return a.replicas
}
