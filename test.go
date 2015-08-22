package main

import (
	"fmt"
	"github.com/nictuku/dht"
	"os"
	"time"
)

func main() {
	// randomly generated hash: AE6D4306F4AE6D4306F4AE6D4306F4AE6D4306F4
	// ubuntu-12.04.4-desktop-amd64.iso: deca7a89a1dbdc4b213de1c0d5351e92582f31fb
	ih, err := dht.DecodeInfoHash("AE6D4306F4AE6D4306F4AE6D4306F4AE6D4306F4")
	if err != nil {
		fmt.Fprintf(os.Stderr, "DecodeInfoHash error: %v\n", err)
		os.Exit(1)
	}
	// Starts a DHT node with the default options. It picks a random UDP port. To change this, see dht.NewConfig.
	d, err := dht.New(nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "New DHT error: %v", err)
		os.Exit(1)
	}
	go d.Run()
	go drain(d)

	for {
		d.PeersRequest(string(ih), false)
		time.Sleep(5 * time.Second)
	}
}

func drain(n *dht.DHT) {
	count := 0
	fmt.Println("=========================== DHT")
	fmt.Println("Note that there are many bad nodes that reply to anything you ask.")
	fmt.Println("Peers found:")
	for r := range n.PeersRequestResults {
		for _, peers := range r {
			for _, x := range peers {
				fmt.Printf("%d: %v\n", count, dht.DecodePeerAddress(x))
				count++
			}
		}
	}
}
