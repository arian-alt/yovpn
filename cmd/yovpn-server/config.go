package main

import (
	"flag"
	"os"
)

type serverConfig struct {
	Port  string
	Token string
}

var portFlag = flag.String("port", "8080", "Port to run HTTP server on")
var tokenFlag = flag.String("token", "", "DigitalOcean access token")

func createConfig() *serverConfig {
	flag.Parse()

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = *portFlag
	}

	token := os.Getenv("DIGITALOCEAN_TOKEN")
	if len(token) == 0 {
		token = *tokenFlag
	}

	return &serverConfig{
		Port:  port,
		Token: token,
	}
}
