package main

import (
	"github.com/michey/gokkan/data"
	"github.com/michey/gokkan/driver"
	"log"
)

func main() {
	c, err := driver.CANConnect("/dev/ttyUSB0", 115200)
	if err != nil {
		log.Fatal(err)
	}
	frame := data.SimpleCANFrameConstruct(0xff, 7, []byte{0, 1, 2, 3, 4, 5, 6, 7})
	err = c.SendFrame(frame)
	if err != nil {
		log.Fatal(err)
	}
}
