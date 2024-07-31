// PLEASE NOTE THAT THIS DOES NOT WORK! WE STILL NEED TO PROPERLY SETUP TINYGO WITH OUR BOARD AND CPU CONFIG!

package main

import (
	"machine"

	"tinygo.org/x/drivers/i2csoft"
)

var screenI2C = i2csoft.New(machine.Pin(4), machine.Pin(5))

func InitDisplay() {
	screenI2C.Configure(i2csoft.I2CConfig{Frequency: 100 * machine.KHz, SCL: machine.Pin(4), SDA: machine.Pin(5)})
}
