package app

import "testing"

func TestNewAutoScale(t *testing.T) {
	var expectReplicas int32 = 1
	scale := NewAutoScale(expectReplicas)
	if replicas := scale.GetReplicas(); replicas != expectReplicas {
		t.Errorf("Incorrect Replicas: %d, get %d", expectReplicas, replicas)
	}
}
