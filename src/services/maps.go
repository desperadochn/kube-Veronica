package services

import (
	"fmt"
	"k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"reflect"
	"sync"
)

type DeploymentMap struct {
	data sync.Map
}

func (this *DeploymentMap)Add(dep *v1.Deployment)  {
	if list,ok:=this.data.Load(dep.Namespace);ok {
		list=append(list.([]*v1.Deployment),dep)
		this.data.Store(dep.Namespace,list)

	}else {
		this.data.Store(dep.Namespace,[]*v1.Deployment{dep})
	}
}

func (this *DeploymentMap)Update(dep *v1.Deployment) error {
	if list,ok:=this.data.Load(dep.Namespace);ok {
		for i,range_dep:=range list.([]*v1.Deployment){
			if range_dep.Name==dep.Name{
				list.([]*v1.Deployment)[i]=dep
			}
		}
	}
	return nil
}

func (this *DeploymentMap)Delete(dep *v1.Deployment)  {
	if list,ok:=this.data.Load(dep.Namespace);ok{
		for i,range_dep:=range list.([]*v1.Deployment){
			if range_dep.Name==dep.Name{
				NewList:= append(list.([]*v1.Deployment)[:i], list.([]*v1.Deployment)[i+1:]...)
				this.data.Store(dep.Namespace,NewList)
				break
			}
		}
	}

}

func (this *DeploymentMap)ListByNS(ns string) ([]*v1.Deployment,error) {
	if list,ok:= this.data.Load(ns);ok {
		return list.([]*v1.Deployment),nil
	}
	return nil,fmt.Errorf("record not found")
}

func (this *DeploymentMap)GetDeployment(ns string,depname string) (*v1.Deployment,error)  {
	if list,ok:=this.data.Load(ns);ok {
		for _,item:=range list.([]*v1.Deployment){
			if item.Name==depname{
				return item,nil
			}
		}
	}
	return nil,fmt.Errorf("record not found")
}

type PodMapStruct struct {
	data sync.Map
}

func (this *PodMapStruct)ListByNs(ns string) []*corev1.Pod  {
	if list,ok:=this.data.Load(ns);ok {
		return list.([]*corev1.Pod)
	}
	return nil
}
func (this *PodMapStruct)Get(ns string,podName string) *corev1.Pod {
	if list,ok:=this.data.Load(ns);ok {
		for _,pod :=range list.([]*corev1.Pod) {
			if pod.Name==podName{
				return pod
			}

		}
	}
	return nil
}
func (this *PodMapStruct)Add(pod *corev1.Pod)  {
	if list,ok:=this.data.Load(pod.Namespace);ok {
		list=append(list.([]*corev1.Pod),pod)
		this.data.Store(pod.Namespace,list)
	}else {
		this.data.Store(pod.Namespace,[]*corev1.Pod{pod})
	}
}

func (this *PodMapStruct)Update(pod *corev1.Pod) error {
	if list,ok:=this.data.Load(pod.Namespace);ok {
		for i,range_pod:= range list.([]*corev1.Pod) {
			if range_pod.Name==pod.Name {
				list.([]*corev1.Pod)[i]=pod
			}
		}
		return nil
	}
	return fmt.Errorf("Pod-%s not found",pod.Name)
}

func (this *PodMapStruct)Delete(pod *corev1.Pod)  {
	if list,ok:=this.data.Load(pod.Namespace);ok {
		for i,range_pod:= range list.([]*corev1.Pod) {
			if range_pod.Name==pod.Name {
				NewList:= append(list.([]*corev1.Pod)[:i], list.([]*corev1.Pod)[:i+1]...)
				this.data.Store(pod.Namespace,NewList)
				break
			}
		}
	}
}

func (this *PodMapStruct)ListByLabels(ns string,labels []map[string]string) ([]*corev1.Pod,error) {
	ret:=make([]*corev1.Pod,0)
	if list,ok:=this.data.Load(ns);ok {
		for _,pod:=range list.([]*corev1.Pod){
			for _,label:= range labels{
				if reflect.DeepEqual(pod.Labels,label) {
					ret=append(ret,pod)
				}
			}
		}
		return ret,nil
	}
	return nil,fmt.Errorf("pods not found ")
}

func (this *PodMapStruct)DEBUG_ListByNS(ns string) ([]*corev1.Pod)  {
	ret:=make([]*corev1.Pod,0)
	if list,ok:=this.data.Load(ns);ok{
		for _,pod:=range list.([]*corev1.Pod){
			ret=append(ret,pod)
		}
	}
	return ret
}