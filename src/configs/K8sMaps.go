package configs

import (
	"k8s.io/client-go/kubernetes"
	"kube-Veronica/src/core"
)

type K8sMap struct {
	K8sClient *kubernetes.Clientset
}

func NewK8sMaps() *K8sMap  {
	return &K8sMap{}
}

func (this *K8sMap)InitDepMap() *core.DeploymentMap {
	return &core.DeploymentMap{}
}

func (this *K8sMap)InitPodMap() *core.PodMapStruct {
	return &core.PodMapStruct{}

}