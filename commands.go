package main

import (
	"fmt"
	"log"

	"github.com/fstanis/go-hive/hive"
)

type commandFunc func(*hive.Client)

var command = map[string]commandFunc{
	"list":                commandList,
	"lights":              commandLights,
	"turnon":              commandTurnOn,
	"turnoff":             commandTurnOff,
	"toggle":              commandToggle,
	"setbrightness":       commandSetBrightness,
	"sethsv":              commandSetHSV,
	"setcolortemperature": commandSetColorTemperature,
	"motion":              commandMotion,
}

func commandList(client *hive.Client) {
	fmt.Println("Listing all devices...\n")
	for _, d := range client.Devices() {
		fmt.Printf("ID:\t%s\nName:\t%s\nType:\t%s\n\n", d.ID(), d.Name(), d.Type())
	}
}

func commandLights(client *hive.Client) {
	fmt.Println("Listing lights...\n")
	for _, d := range client.Devices() {
		if d.IsLight() {
			fmt.Printf("ID:\t\t%s\nName:\t\t%s\nIs colored:\t%v\nIs on:\t\t%v\n\n", d.ID(), d.Name(), d.IsColorLight(), d.IsOn())
		}
	}
}

func commandTurnOn(client *hive.Client) {
	device := mustBeLight(deviceById(client))
	if err := device.Do(hive.NewChange().TurnOn()); err != nil {
		log.Fatalf("Failed to turn on light: %v", err)
	}
}

func commandTurnOff(client *hive.Client) {
	device := mustBeLight(deviceById(client))
	if err := device.Do(hive.NewChange().TurnOff()); err != nil {
		log.Fatalf("Failed to turn off light: %v", err)
	}
}

func commandToggle(client *hive.Client) {
	device := mustBeLight(deviceById(client))

	var err error
	if device.IsOn() {
		err = device.Do(hive.NewChange().TurnOff())
	} else {
		err = device.Do(hive.NewChange().TurnOn())
	}
	if err != nil {
		log.Fatalf("Failed to toggle light: %v", err)
	}
}

func commandSetBrightness(client *hive.Client) {
	device := mustBeLight(deviceById(client))
	brightness := intArgument(2)
	if err := device.Do(hive.NewChange().Brightness(brightness)); err != nil {
		log.Fatalf("Failed to set brightness: %v", err)
	}
}

func commandSetHSV(client *hive.Client) {
	device := mustBeColorLight(deviceById(client))
	colorArg := parseColorArgument(strArgument(2))
	if err := device.Do(hive.NewChange().Color(colorArg)); err != nil {
		log.Fatalf("Failed to set color temperature: %v", err)
	}
}

func commandSetColorTemperature(client *hive.Client) {
	device := mustBeColorLight(deviceById(client))
	temperature := intArgument(2)
	if err := device.Do(hive.NewChange().ColorTemperature(temperature)); err != nil {
		log.Fatalf("Failed to set color temperature: %v", err)
	}
}

func commandMotion(client *hive.Client) {
	device := mustBeMotionSensor(deviceById(client))
	hasMotion := device.HasMotion()
	if hasMotion {
		fmt.Println("The sensor is currently seeing motion.\n")
	} else {
		fmt.Println("The sensor is NOT currently seeing motion.\n")
	}
	fmt.Printf("Last motion started:\t%s\nLast motion stopped:\t%s\n", device.LastMotionEnd(), device.LastMotionStart())
}
