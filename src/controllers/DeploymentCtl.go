package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	"k8s.io/client-go/kubernetes"
	"kube-Veronica/src/services"
)

type DeploymentCtl struct {
	K8sClient *kubernetes.Clientset  `inject:"-"`
	DepService *services.DeploymentService  `inject:"-"`
}

func NewDeploymentCtl() *DeploymentCtl  {
	return &DeploymentCtl{}
}

func (this *DeploymentCtl)GetList(c *gin.Context) goft.Json  {
	ns:=c.DefaultQuery("ns","default")
	return gin.H{
		"code":20000,
		"data":this.DepService.ListAll(ns),
	}

}

func (this *DeploymentCtl)Build(goft *goft.Goft)  {
	goft.Handle("GET","/deployments",this.GetList)
	
}

func (this *DeploymentCtl)Name() string  {
	return "DeploymentCtl"
}