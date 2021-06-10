package services

import (
	"github.com/shenyisyn/goft-gin/goft"
	"k8s.io/api/apps/v1"
	"kube-Veronica/src/core"
	"kube-Veronica/src/models"
)

//@Service
type DeploymentService struct {
	DepMap *core.DeploymentMap  `inject:"-"`
	Common *CommonService `inject:"-"`
}

func NewDeploymentService() *DeploymentService {
	return &DeploymentService{}
}
func (*DeploymentService)getDeploymentCondition(dep *v1.Deployment) string {
	for _,item:=range dep.Status.Conditions {
		if string(item.Type)=="Available" && string(item.Status)!="True" {
			return item.Message
		}
	}
	return ""
}

func (*DeploymentService)getDeploymentIsComplete(dep *v1.Deployment) bool {
	return dep.Status.Replicas==dep.Status.AvailableReplicas
}

func (this *DeploymentService)ListAll(namespace string) (ret []*models.Deployment)  {
	depList,err:=this.DepMap.ListByNS(namespace)
	goft.Error(err)
	for _,item:=range depList{
		ret=append(ret,&models.Deployment{Name: item.Name,
			NameSpaces: item.Name,
			Replicas:   [3]int32{item.Status.Replicas, item.Status.AvailableReplicas, item.Status.UnavailableReplicas},
			Images:     this.Common.GetImages(*item),
			IsComplete: this.getDeploymentIsComplete(item),
			Message:    this.getDeploymentCondition(item),
		})

	}
	return
}