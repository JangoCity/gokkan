package driver

import (
	"github.com/jacobsa/go-serial/serial"
	"github.com/michey/gokkan/protocol"
	"io"
)

type SerialCANConnection struct {
	port          *io.ReadWriteCloser
	protocol      *protocol.IProtocol
	readEnabled   bool
	readSkip      bool
	writeEnabled  bool
	inputChannel  chan messages.ToDevice
	outputChannel chan messages.FromDevice
}

func CANConnectWithPotato(serialName *string, baudRate int) (connection SerialCANConnection, err error) {
	p := protocol.InitPotatoProtocol()
	return CANConnect(serialName, baudRate, p)
}

func CANConnect(serialName *string, baudRate int, iProtocol protocol.IProtocol) (connection SerialCANConnection, err error) {
	options := serial.OpenOptions{
		BaudRate:        1500000,
		PortName:        *serialName,
		DataBits:        8,
		StopBits:        1,
		MinimumReadSize: 4}

	s, err := serial.Open(options)
	if err != nil {
		return SerialCANConnection{nil, nil, true, true, true, nil, nil}, err
	}
	input := make(chan messages.ToDevice, 256)
	output := make(chan messages.FromDevice, 256)

	c := SerialCANConnection{&s, &iProtocol, true, true, true, input, output}
	c.Run()
	return c, nil
}

func (conn *SerialCANConnection) Send(frame messages.ToDevice) {
	conn.inputChannel <- frame
}

func (conn *SerialCANConnection) Skip(flag bool) {
	conn.readSkip = flag
}

func (conn *SerialCANConnection) Run() {
	go conn.reader(conn.outputChannel)
	go conn.writer(conn.inputChannel)
}

func (conn *SerialCANConnection) GetChan() <-chan messages.FromDevice {
	return conn.outputChannel
}

func (conn *SerialCANConnection) stop() {
	conn.readEnabled = false
	conn.writeEnabled = false
	(*conn.protocol).Stop()
}

func (conn *SerialCANConnection) reader(output chan<- messages.FromDevice) {
	p := *conn.protocol
	bytes := make(chan byte, 1)
	b := make([]byte, 32)

	go p.Decode(bytes, output)

	for conn.readEnabled {
		n, _ := (*conn.port).Read(b)
		if n > 0 {
			for i := 0; i < n; i++ {
				bytes <- b[i]
			}
		}
	}
}

func (conn *SerialCANConnection) writer(input <-chan messages.ToDevice) {
	for conn.writeEnabled {
		f := <-input
		p := *conn.protocol
		b := p.Encode(&f)
		(*conn.port).Write(b)
	}
}
