package wscore

type WsMessage struct {
	MessageType int
	MessageDate []byte
}

func NewWsMessage(messageType int, messageData []byte) *WsMessage {
	return &WsMessage{MessageType: messageType,MessageDate: messageData}
}