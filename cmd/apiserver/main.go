package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/adametsderschopfer/http-rest-api_go-edu/internal/app/apiserver"
	"log"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "Path to config file")
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	if err := apiserver.Bootstrap(config); err != nil {
		log.Fatal(err)
	}
}
