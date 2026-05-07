package main

const (
	CmdAuthentic byte = 0x00
	CmdConnect   byte = 0x01
	CmdAssociate byte = 0x02
	CmdHeartbeat byte = 0x03
)

const (
	RespSuccess byte = 0x01
	RespFailed  byte = 0x00
)

type TUICPacket struct {
	Command byte
	UUID    [16]byte
	Payload []byte
}
