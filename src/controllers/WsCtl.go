package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	"log"
	"kube-Veronica/src/wscore"
)

type WsCtl struct {}

func NewWsCtl() *WsCtl {
	return &WsCtl{}
}


func (this *WsCtl)Connect(c *gin.Context) string {
	client,err:=wscore.Upgrader.Upgrade(c.Writer,c.Request,nil)
	if err!=nil{
		log.Println(err)
		return err.Error()
	}else {
		wscore.ClientMap.Store(client)
		return "succss"
	}

}
func (this *WsCtl)Build(goft *goft.Goft)  {
	goft.Handle("GET","/ws",this.Connect)
}

func (this *WsCtl)Name() string  {
	return "WsCtl"
}