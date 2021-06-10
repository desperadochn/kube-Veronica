package models

type Pod struct {
	Name string
	NameSpace string
	Images string
	NodeName string
	Ip []string
	Phase string
	IsReady bool
	Message string
	CreateTime string

}
