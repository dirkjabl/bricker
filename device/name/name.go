// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// A simple function for converting the device identifer to a real name.
package name

var devicenames = map[uint16]string{
	11:  "Brick DC",
	13:  "Brick Master",
	14:  "Brick Servo",
	15:  "Brick Stepper",
	16:  "Brick IMU",
	21:  "Bricklet Ambient Light",
	23:  "Bricklet Current12",
	24:  "Bricklet Current25",
	25:  "Bricklet Distance IR",
	26:  "Bricklet Dual Relay",
	27:  "Bricklet Humidity",
	28:  "Bricklet IO-16",
	29:  "Bricklet IO-4",
	210: "Bricklet Joystick",
	211: "Bricklet LCD 16x2",
	212: "Bricklet LCD 20x4",
	213: "Bricklet Linear Poti",
	214: "Bricklet Piezo Buzzer",
	215: "Bricklet Rotary Poti",
	216: "Bricklet Temperature",
	217: "Bricklet Temperature IR",
	218: "Bricklet Voltage",
	219: "Bricklet Analog In",
	220: "Bricklet Analog Out",
	221: "Bricklet Barometer",
	222: "Bricklet GPS",
	223: "Bricklet Industrial Digital In 4",
	224: "Bricklet Industrial Digital Out 4",
	225: "Bricklet Industrial Quad Relay",
	226: "Bricklet PTC",
	227: "Bricklet Voltage/Current",
	228: "Bricklet Industrial Dual 0-20mA",
	229: "Bricklet Distance US",
	230: "Bricklet Dual Button",
	231: "Bricklet LED Strip",
	232: "Bricklet Moisture",
	233: "Bricklet Motion Detector",
	234: "Bricklet Multi Touch",
	235: "Bricklet Remote Switch",
	236: "Bricklet Rotary Encoder",
	237: "Bricklet Segment Display 4x7",
	238: "Bricklet Sound Intensity",
	239: "Bricklet Tilt",
	240: "Bricklet Hall Effect",
	241: "Bricklet Line",
	242: "Bricklet Piezo Speaker",
	243: "Bricklet Color",
	245: "Bricklet Heart Rate",
	246: "Bricklet NFC/RFID"}

// Name converts a device identifer to a string
func Name(id uint16) string {
	if name, ok := devicenames[id]; ok {
		return name
	}
	return "unknown hardware"
}
