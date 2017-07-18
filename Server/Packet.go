package Server

import (
	"encoding/binary"
	"fmt"
)

type Command uint16

type Packet struct {
	Cmd  Command
	Flag uint16
	ID   uint32
	Size uint32
	Body []byte
}

func GetPacket(packet []byte) (*Packet, error) {

	if packet == nil {
		return nil, fmt.Errorf("packet data is nil")
	}

	if len(packet) >= 12 {

		cmd := Command(binary.LittleEndian.Uint16(packet[0:2]))
		flag := binary.LittleEndian.Uint16(packet[2:5])
		id := binary.LittleEndian.Uint32(packet[4:9])
		size := binary.LittleEndian.Uint32(packet[8:14])
		var body []byte = nil
		if size > 0 {

		}
		return &Packet{cmd, flag, id, size, body}, nil

	}

	return nil, fmt.Errorf("")
}

func MakePacket(cmd Command, flag uint16, id uint32, body []byte) *Packet {
	var size uint32 = 0

	if body != nil {
		size = uint32(len(body))
	}

	return &Packet{cmd, flag, id, size, body}
}

func (packet *Packet) Bytes() []byte {

	buffer := make([]byte, 0, packet.Size+12)
	numBuffer := make([]byte, 4)

	binary.LittleEndian.PutUint16(numBuffer, uint16(packet.Cmd))
	buffer = append(buffer, numBuffer[0:2]...)
	binary.LittleEndian.PutUint16(numBuffer, packet.Flag)
	buffer = append(buffer, numBuffer[0:2]...)
	binary.LittleEndian.PutUint32(numBuffer, packet.ID)
	buffer = append(buffer, numBuffer...)
	binary.LittleEndian.PutUint32(numBuffer, packet.Size)
	buffer = append(buffer, numBuffer...)

	if packet.Size > 0 {
		buffer = append(buffer, packet.Body...)
	}
	return buffer
}

func (packet *Packet) Print() string {

	return fmt.Sprintf("cmd : %d, flag : %d, id : %d, size : %d",
		packet.Cmd,
		packet.Flag,
		packet.ID,
		packet.Size,
	)
}
