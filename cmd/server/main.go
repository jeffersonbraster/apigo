package main

import "github.com/jeffersonbraster/apigo/configs"

func main() {
		config, _ := configs.LoadConfig(".")
		println(config.DBDrive)
}