package main

import (
	"github.com/vishvananda/netlink"
	"log"
)

var (
	iface = "enp2s0"
)

func main() {
	log.Println("Start of application")

	link, err := netlink.LinkByName(iface)
	if err != nil {
		log.Fatalf("netlink.LinkByName %v: %v\n", iface, err)
	}

	linkIndex := link.Attrs().Index

	done := make(chan struct{})
	// defer close(done)

	addrChannel := make(chan netlink.AddrUpdate)
	defer close(addrChannel)

	if err := netlink.AddrSubscribe(addrChannel, done); err != nil {
		log.Fatalf("netlink.AddrSubscribe: %v", err)
	}

	for addr := range addrChannel {
		if addr.LinkIndex == linkIndex {
			log.Printf("%+v\n", addr)
		}
	}

	close(done)
}
