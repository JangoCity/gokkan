package driver

import (
	"fmt"
	"github.com/michey/gokkan/data"
	"github.com/michey/gokkan/protocol"
	"github.com/tarm/serial"
	"log"
	"time"
)

type SerialCANConnection struct {
	port          *serial.Port
	protocol      *protocol.IProtocol
	readEnabled   bool
	writeEnabled  bool
	inputChannel  chan data.CANFrame
	outputChannel chan data.Response
}

func CANConnectWithPotato(serialName string, baudRate int) (connection SerialCANConnection, err error) {
	p := protocol.PotatoProtocol{}
	return CANConnect(serialName, baudRate, &p)
}

func CANConnect(serialName string, baudRate int, iProtocol protocol.IProtocol) (connection SerialCANConnection, err error) {
	s, err := serial.OpenPort(&serial.Config{Name: serialName, Baud: baudRate, ReadTimeout: time.Second * 5})
	if err != nil {
		return SerialCANConnection{nil, nil, true, true, nil, nil}, err
	}
	input := make(chan data.CANFrame, 256)
	output := make(chan data.Response, 256)

	c := SerialCANConnection{s, &iProtocol, true, true, input, output}
	c.Run()
	return c, nil
}

func (conn *SerialCANConnection) Send(frame data.CANFrame) {
	conn.inputChannel <- frame
}

func (conn *SerialCANConnection) Run() {
	go conn.reader(conn.outputChannel)
	go conn.writer(conn.inputChannel)
}

func (conn *SerialCANConnection) GetChan() chan data.Response {
	return conn.outputChannel
}

func (conn *SerialCANConnection) stop() {
	conn.readEnabled = false
	conn.writeEnabled = false
}

func (conn *SerialCANConnection) reader(output chan data.Response) {
	d := make([]byte, 1024)
	dataPosition := 0
	p := *conn.protocol

	b := make([]byte, 256)

	for conn.readEnabled {
		n, _ := conn.port.Read(b)
		fmt.Println(n)
		if n > 0 {
			for i := 0; i < n; i++ {
				d[dataPosition] = b[i]

				fmt.Println(d[dataPosition])
				if d[dataPosition] == '\n' {
					validData := make([]byte, dataPosition+1)
					copy(validData, d)
					fmt.Println(validData)
					err, frame := p.Decode(validData)
					if err != nil {
						log.Fatal(err, validData)
					} else {
						response := data.Response{frame, time.Now().UnixNano() / int64(time.Millisecond)}
						output <- response
					}
					dataPosition = 0
				}

				dataPosition++
				if dataPosition >= 1024 {
					dataPosition = 0
				}
			}
		}
	}
}

func (conn *SerialCANConnection) writer(input chan data.CANFrame) {
	f := <-input
	fmt.Println("lol, getmsg")
	p := *conn.protocol
	b := p.Encode(&f)
	fmt.Println("%x", b)
	conn.port.Write(b)
}
