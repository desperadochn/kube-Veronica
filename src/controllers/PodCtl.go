package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	"kube-Veronica/src/services"
)

type PodCtl struct {
	PodService *services.PodService  `inject:"-"`
}

func NewPodCtl() *PodCtl {
	return &PodCtl{}
}

func (this *PodCtl)GetAll(c *gin.Context) goft.Json {
	return this.PodService.ListByNs("zhouchi")
}

func (this *PodCtl)Build(goft *goft.Goft)  {
	goft.Handle("GET","/Pods",this.GetAll)
}

func (this *PodCtl)Name() string {
	return "PodCtl"
}