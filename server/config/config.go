package config

import (
	"encoding/json"
	"flag"
	"io/ioutil"
)

type Config struct {
	ArduinoSerialName string `json:"arduino_serial_name"`
	ArduinoSerialBaud int    `json:"arduino_serial_baud"`
}

var config *Config

func GetConfig() *Config {
	if config == nil {
		config = new(Config).readConfig()
	}
	return config
}

func (c Config) readConfig() *Config {
	flagName := c.readFlags()
	file, err := ioutil.ReadFile(*flagName)
	err = json.Unmarshal(file, &config)
	if err != nil {
		return nil
	}
	return config
}

func (c Config) readFlags() (configFile *string) {
	configFile = flag.String("config", "serial-config.json", "Path to config file")
	flag.Parse()
	return
}
