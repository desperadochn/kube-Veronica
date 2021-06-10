package core

import (
	v1 "k8s.io/api/apps/v1"
	"log"
	corev1 "k8s.io/api/core/v1"
)

type DepHandler struct {
	DepMap *DeploymentMap `inject:"-"`
}



func (this *DepHandler)OnAdd(obj interface{})  {
	this.DepMap.Add(obj.(*v1.Deployment))
}
func (this *DepHandler)OnUpdate(oldObj,newObj interface{})  {
	err:=this.DepMap.Update(newObj.(*v1.Deployment))
	if err !=nil {
		log.Panicln(err)
	}
}
func (this *DepHandler)OnDelete(obj interface{})  {
	if d,ok:=obj.(*v1.Deployment);ok {
		this.DepMap.Delete(d)
	}
}

type PodHandler struct {
	PodMap *PodMapStruct `inject:"-"`
}

func (this *PodHandler)OnAdd(obj interface{})  {
	this.PodMap.Add(obj.(*corev1.Pod))
}
func (this *PodHandler)OnUpdate(oldObj,newObj interface{})  {
	err:=this.PodMap.Update(newObj.(*corev1.Pod))
	if err !=nil {
		log.Println(err)
	}
}

func (this *PodHandler)OnDelete(obj interface{})  {
	if d,ok:=obj.(*corev1.Pod);ok {
		this.PodMap.Delete(d)
	}
}