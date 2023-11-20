package main

import (
	"fmt"
	"os"

	"github.com/songgao/water"
)

func main() {
	iface, err := water.New(water.Config{
		DeviceType: water.TUN,
	})
	if err != nil {
		fmt.Printf("error creating TUN interface: %v", err)
		os.Exit(1)
	}
	defer iface.Close()

	buf := make([]uint8, 1504)
	for {
		n, err := iface.Read(buf)
		if err != nil {
			fmt.Printf("error reading from interface: %v\n", err)
		}
		fmt.Printf("%d bytes read\n", n)
	}
}
