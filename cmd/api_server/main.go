package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/Eclipsium/go_rest_api/internal/api_server"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/api_server.toml", "path to config file")
}
func main() {
	flag.Parse()

	config := api_server.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	s := api_server.New(config)
	if error := s.Start(); error != nil {
		log.Fatal(error)
	}
}
