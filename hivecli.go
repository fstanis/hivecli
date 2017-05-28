package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/fstanis/hivecli/config"
)

var (
	configPath = flag.String("config", "", "location of the file where the config data is saved")
)

const usage = `Usage: hivecli [OPTION]... COMMAND [ARGUMENT]...
Runs the given command using the Hive API and displays the result, if any.

Options:
  --config=FILE              specifies the location of the file to store the
                             configuration in.

Commands:
  list                             lists all devices present
  lights                           lists all lights present
  turnon <id>                      turns the given light on
  turnoff <id>                     turns the given light off
  toggle <id>                      toggles the given light on or off
  setbrightness <id> <value>       sets the brightness of the given light
  sethsv <id> <h>,<s>,<v>          sets the color of the given light to the
                                     specified HSV (hue, saturation and value).
  setcolortemperature <id> <value> sets the color temperature of the given light
  motion <id>                      reads the status of the given motion sensor
`

func printUsage() {
	fmt.Println(usage)
}

func main() {
	flag.Parse()
	cmd := flag.Arg(0)
	if cmd == "" {
		fmt.Println("No command specified.\n")
		printUsage()
		os.Exit(2)
	}
	if _, ok := command[cmd]; !ok {
		fmt.Printf("Command %q not found.\n\n", cmd)
		printUsage()
		os.Exit(2)
	}

	conf, _ := config.FromFile(*configPath)
	client, err := connect(conf)
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Error trying to run %s: %v\n", cmd, r)
		}
	}()
	commandFunc := command[cmd]
	commandFunc(client)
}
