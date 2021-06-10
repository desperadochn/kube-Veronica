package configs

import (
	"kube-Veronica/src/core"
)

type K8sHandler struct {

}
func NewK8sHandler() *K8sHandler {
	return &K8sHandler{}
}
func (this *K8sHandler) DepHandlers() *core.DepHandler {
	return &core.DepHandler{}
}

func (this *K8sHandler)PodHandlers() *core.PodHandler  {
	return &core.PodHandler{}
}