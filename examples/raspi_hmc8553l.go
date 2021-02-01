// +build example
//
// Do not build by default.

/*
 How to run

	go run examples/firmata_hmc8553l.go
*/

package main

import (
	"fmt"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/i2c"
	"gobot.io/x/gobot/platforms/raspi"
)

func main() {
	raspi := raspi.NewAdaptor()
	hmc8553l := i2c.NewHMC8553LDriver(raspi)

	work := func() {
		gobot.Every(100*time.Millisecond, func() {
			heading, _ := hmc8553l.Heading()
			fmt.Println("Heading", heading)
		})
	}

	robot := gobot.NewRobot("hmc8553LBot",
		[]gobot.Connection{raspi},
		[]gobot.Device{hmc8553l},
		work,
	)

	robot.Start()
}
