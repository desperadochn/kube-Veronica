package models

type Deployment struct {
	Name string
	NameSpaces string
	Replicas [3]int32
	Images string
	IsComplete bool
	Message string
	CreateTime string
	Pods []*Pod
}