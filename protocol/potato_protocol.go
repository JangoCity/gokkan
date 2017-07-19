package protocol

import (
	"github.com/golang/protobuf/proto"
	"github.com/michey/gokkan/messages"
	"log"
)

type PotatoProtocol struct {
	IProtocol
	buffer         []byte
	bufferPosition uint16
	read           bool
}

func InitPotatoProtocol() *PotatoProtocol {
	buffer := make([]byte, 512)
	return &PotatoProtocol{buffer: buffer, bufferPosition: 0, read: true}
}

func (p *PotatoProtocol) Encode(frame *messages.ToDevice) []byte {
	data, err := proto.Marshal(frame)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	return data
}

func (p *PotatoProtocol) Stop() {
	p.read = false
}

func (p *PotatoProtocol) Decode(bytes <-chan byte, responses chan<- messages.FromDevice) {
	msg := &messages.FromDevice{}
	for p.read {
		byte := <-bytes
		p.buffer[p.bufferPosition] = byte
		p.bufferPosition++
		err := proto.Unmarshal(p.buffer, msg)
		if err == nil {
			responses <- *msg
			p.buffer = make([]byte, 512)
			p.bufferPosition = 0
		}

	}
}
