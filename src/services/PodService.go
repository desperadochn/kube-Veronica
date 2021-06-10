package services

import (
	"kube-Veronica/src/core"
)

type PodService struct {
	PodMap *core.PodMapStruct `inject:"-"`
}

func NewPodService() *PodService  {
	return &PodService{}
}

func (this *PodService)ListByNs(ns string) interface{} {
	return this.PodMap.ListByNs(ns)
}
