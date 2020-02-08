// https://www.socketloop.com/references/golang-net-interfaces-function-examples
package main

import (
	"fmt"
	"net"
)

// GetInterface get interface info.
func GetInterface() {

	interfaces, err := net.Interfaces()

	if err != nil {
		fmt.Print(err)
		return
	}

	for _, i := range interfaces {

		fmt.Printf("Name : %v \n", i.Name)

		// see http://golang.org/pkg/net/#Flags
		fmt.Println("Interface type and supports : ", i.Flags.String())
	}

}
