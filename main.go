package main

import (
	"fmt"

	"flexera.com/config"
	"flexera.com/handler"
)

func main() {
	conf, err := config.LoadConfig("./config/")
	if err != nil {
		fmt.Println("Error in loading configurations!")
	}

	err = handler.ProcessCSV(conf.LargeCSVURI)
	if err != nil {
		fmt.Println("Error::", err)
	}

}
