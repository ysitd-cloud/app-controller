package kubernetes

import (
	"fmt"

	"github.com/ysitd-cloud/app-controller/app"
	"github.com/ysitd-cloud/app-controller/version"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (d *deploymentV1) GetApplication() app.Application {
	return d.application
}

func (d *deploymentV1) getName() string {
	return fmt.Sprintf("app-%s", d.application.GetID())
}

func (d *deploymentV1) getLabels() map[string]string {
	labels := make(map[string]string)
	labels["deployerGitCommit"] = version.Version.GitCommit
	labels["deployerGitVersion"] = version.Version.GoVersion
	labels["deployerGoVersion"] = version.Version.GoVersion
	labels["app"] = d.application.GetID()
	return labels
}

func (d *deploymentV1) getObjectMeta() metav1.ObjectMeta {
	return metav1.ObjectMeta{
		Name:   d.getName(),
		Labels: d.getLabels(),
	}
}
