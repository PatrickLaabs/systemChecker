package main

import (
	configloader "systemChecker/pkg/configloader"
	configreloader "systemChecker/pkg/configreloader"
	server "systemChecker/pkg/server"
)

func main() {
	configloader.Configloader()
	server.Server()
	configreloader.Configreloader()
}
