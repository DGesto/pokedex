package main

import (
	"fmt"
)

func commandMapB(conf *Config) error {
	locationAreas := getLocationAreas(conf.Previous, conf)
	for _, locationArea := range locationAreas {
		fmt.Print(locationArea.Name + "\n")
	}
	return nil
}
