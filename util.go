package main

import (
	"flag"
	"log"
	"strconv"
	"strings"

	"github.com/fstanis/go-hive/hive"
)

func deviceById(client *hive.Client) *hive.Device {
	id := flag.Arg(1)
	if id == "" {
		log.Fatalf("Must specify ID as parameter")
	}
	device := client.Device(id)
	if device == nil {
		log.Fatalf("Device with the given ID not found")
	}
	return device
}

func intArgument(i int) int {
	argStr := flag.Arg(i)
	if argStr == "" {
		log.Fatalf("Numeric argument expected")
	}
	arg, err := strconv.Atoi(argStr)
	if err != nil {
		log.Fatalf("Numeric argument expected, got %q", argStr)
	}
	return arg
}

func strArgument(i int) string {
	argStr := flag.Arg(i)
	if argStr == "" {
		log.Fatalf("Argument expected")
	}
	return argStr
}

func parseColorArgument(arg string) (hsv hive.HSV) {
	colors := strings.Split(arg, ",")
	if len(colors) != 3 {
		log.Fatalf("Expected 3 comma-separated numbers as argument")
	}

	var err error
	if hsv.Hue, err = strconv.Atoi(colors[0]); err != nil {
		log.Fatalf("Expected 3 comma-separated numbers as argument")
	}
	if hsv.Saturation, err = strconv.Atoi(colors[1]); err != nil {
		log.Fatalf("Expected 3 comma-separated numbers as argument")
	}
	if hsv.Value, err = strconv.Atoi(colors[2]); err != nil {
		log.Fatalf("Expected 3 comma-separated numbers as argument")
	}
	return hsv
}

func mustBeLight(d *hive.Device) *hive.Device {
	if !d.IsLight() {
		log.Fatalf("Device with the given ID is not a light")
	}
	return d
}

func mustBeColorLight(d *hive.Device) *hive.Device {
	if !d.IsColorLight() {
		log.Fatalf("Device with the given ID is not a color light")
	}
	return d
}

func mustBeMotionSensor(d *hive.Device) *hive.Device {
	if !d.IsMotionSensor() {
		log.Fatalf("Device with the given ID is not a motion sensor")
	}
	return d
}
