package main

//Dependencies
import (
    /src/internal/logic
	/src/internal/data
)

func main() {
	//var apidata = []northbound{}

	//Read and setup configuration data
	configuration.SetupConfiguration()

	//Lauch Api server
	logic.run()
}
