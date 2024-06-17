package main

//Dependencies
import (
	"doc/src/logic"
	"doc/src/config"
)

func main() {
	//var apidata = []northbound{}

	//Read and setup configuration data
	config.SetupConfiguration()

	//Lauch Api server
	logic.Run()
}
