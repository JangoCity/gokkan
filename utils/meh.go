package main

import (
	"fmt"
	"github.com/michey/gokkan/data"
	"github.com/michey/gokkan/driver"
	"github.com/michey/gokkan/messages"
	"log"
	"time"
)

func main() {
	p := "/dev/ttyUSB0"
	c, err := driver.CANConnectWithPotato(&p, 115200)
	if err != nil {
		log.Fatal(err)
	}
	go pp(c.GetChan())
	baudRate := data.BaudRate{Rate: messages.BaudRate_K1000}
	frame := data.SimpleCANFrameConstruct(0x3f, 8, []byte{0, 1, 2, 3, 4, 5, 6, 7})

	c.Skip(false)
	c.Send(baudRate.ToDevice())
	c.Send(data.GetInit().ToDevice())
	c.Send(frame.ToDevice())
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(10 * time.Second)
}

func pp(c <-chan messages.FromDevice) {
	msg := <-c
	fmt.Println(&msg)
}
