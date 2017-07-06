package main

import (
	"flag"
	"fmt"
	"github.com/michey/gokkan/data"
	"github.com/michey/gokkan/driver"
	"github.com/michey/gokkan/helpers"
	"github.com/michey/gokkan/uds"
	"log"
	"time"
)

func main() {
	portName := flag.String("port", "/dev/ttyUSB0", "Serial port name")

	c, err := driver.CANConnectWithPotato(portName, 115200)
	if err != nil {
		log.Fatal(err)
	}

	r := make(chan data.CANFrame, 10)
	output := make(chan []data.Response, 1)

	go singleFrameSupplier(r)

	f := <-r

	for f.StdId <= data.MAX_STD_ID {
		c.Skip(false)
		helpers.ReadByTimeWindow(3*time.Millisecond, c.GetChan(), output)
		c.Send(f)
		msgs := <-output
		c.Skip(true)
		findCorrectResponse(msgs)
		f = <-r
	}

}

func singleFrameSupplier(c chan data.CANFrame) {
	for i := 0; i <= data.MAX_STD_ID; i++ {
		sf := uds.RequestForSession(uds.DefaultSession)
		c <- *sf.ToFrame(uint32(i))
	}
}

func findCorrectResponse(responses []data.Response) {
	//fmt.Printf("%+v", responses)

	for i := 0; i < len(responses); i++ {
		response := responses[i]
		if response.Data[1] == 0x50 || response.Data[1] == 0x7F {
			fmt.Println(&response)
			fmt.Printf("Well, ECU CAN ID is %#v. \n", response.StdId)
		}
	}
}
