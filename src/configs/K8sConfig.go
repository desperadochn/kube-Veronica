package configs

import (

	"io/ioutil"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog/v2"
	"kube-Veronica/src/core"
	"log"
)

type K8sConfig struct {
	DepHandler *core.DepHandler `inject:"-"`
	PodHandler *core.PodHandler `inject:"-"`
}

func NewK8sConfig() *K8sConfig  {
	return &K8sConfig{}

}

func (*K8sConfig)InitClient() *kubernetes.Clientset{
	var config *rest.Config
	var err error
	var kubeConfig []byte
	if kubeConfig,err= ioutil.ReadFile("/Users/zhouchi/.kube/config"); err != nil {
		klog.Error(err)
		return nil
	}
	if config, err = clientcmd.RESTConfigFromKubeConfig(kubeConfig); err != nil {
		return nil
	}
	c,err:=kubernetes.NewForConfig(config)
	if err!=nil{
		log.Fatal(err)
	}
	return c

}

func (this *K8sConfig)InitInformer() informers.SharedInformerFactory  {
	fact:=informers.NewSharedInformerFactory(this.InitClient(),0)
	depinformer:=fact.Apps().V1().Deployments()
	depinformer.Informer().AddEventHandler(this.DepHandler)
	podinformer:=fact.Core().V1().Pods()
	podinformer.Informer().AddEventHandler(this.PodHandler)
	fact.Start(wait.NeverStop)
	return fact
}
