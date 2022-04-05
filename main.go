package main

import (
	"fmt"

	vmaas "github.com/eepurichandra/go-morpheus-client/internal/pkg/hpegl-vmaas-client"
)

func main() {
	fmt.Println("In Main function")
	msg := vmaas.CallClient()
	fmt.Println(msg)
}
