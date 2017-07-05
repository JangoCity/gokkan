package main

import (
	"flag"
	"fmt"
	"github.com/michey/gokkan/driver"
	"log"
)

func main() {
	portName := flag.String("port", "/dev/ttyUSB0", "Serial port name")

	c, err := driver.CANConnectWithPotato(portName, 115200)
	if err != nil {
		log.Fatal(err)
	}
	c.Skip(false)

	for true {
		msg := <-c.GetChan()
		fmt.Println(&msg)
	}

}
