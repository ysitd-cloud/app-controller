package template

import (
	"fmt"

	"github.com/ysitd-cloud/app-controller/version"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetName(id string) string {
	return fmt.Sprintf("app-%s", id)
}

func getObjectMeta(id string) metav1.ObjectMeta {
	return metav1.ObjectMeta{
		Name:   GetName(id),
		Labels: getLabels(id),
	}
}

func getLabels(id string) map[string]string {
	labels := make(map[string]string)
	labels["deployerGitCommit"] = version.Version.GitCommit
	labels["deployerGitVersion"] = version.Version.GoVersion
	labels["deployerGoVersion"] = version.Version.GoVersion
	labels["app"] = id
	labels["type"] = "userApp"
	return labels
}
