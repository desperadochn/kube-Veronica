package wscore

import (
	"github.com/gorilla/websocket"
	"log"
	"sync"
	"time"
)

type ClientMapStruct struct {
	data sync.Map

}
var ClientMap *ClientMapStruct
func init()  {
	ClientMap=&ClientMapStruct{}
}

func (this *ClientMapStruct)Store(conn *websocket.Conn)  {
	wsClient:=NewWsClient(conn)
	this.data.Store(conn.RemoteAddr().String(),wsClient)
	go wsClient.Ping(time.Second*1)
	go wsClient.ReadLoop()
}

func (this *ClientMapStruct)SendAllDepList(v interface{})  {
	this.data.Range(func(key, value interface{}) bool {
		c:=value.(*WsClient).conn
		    err:=c.WriteJSON(v)
		    if err!=nil {
		    	this.Remove(c)
		    	log.Println(err)
			}
			return true
	})
}

func (this *ClientMapStruct)Remove(conn *websocket.Conn)  {
	this.data.Delete(conn.RemoteAddr().String())
}
