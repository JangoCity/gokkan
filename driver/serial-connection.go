package driver

import (
	"github.com/michey/gokkan/data"
	"github.com/michey/gokkan/protocol"
	"github.com/tarm/serial"
)

type CANConnection struct {
	port *serial.Port
}

func CANConnect(serialName string, baudRate int) (connection CANConnection, err error) {
	s, err := serial.OpenPort(&serial.Config{Name: serialName, Baud: baudRate})
	if err != nil {
		return CANConnection{nil}, err
	}
	return CANConnection{port: s}, nil
}

func (conn *CANConnection) SendFrame(frame data.CANFrame) error {
	_, err := conn.port.Write(protocol.Marshall(frame))
	if err != nil {
		return err
	}
	return nil

}
