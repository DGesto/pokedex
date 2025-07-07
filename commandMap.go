package main

import (
	"fmt"
)

func commandMap(conf *Config) error {
	locationAreas := getLocationAreas(conf.Next, conf)
	for _, locationArea := range locationAreas {
		fmt.Print(locationArea.Name + "\n")
	}
	return nil
}
