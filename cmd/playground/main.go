/*
Copyright 2021 The routerd authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// most of this code comes directly from the example in https://pkg.go.dev/github.com/insomniacslk/dhcp@v0.0.0-20210120172423-cc9239ac6294/dhcpv4/server4

package main

import (
	"log"
	"net"
	"os"

	"github.com/insomniacslk/dhcp/dhcpv4"
	"github.com/insomniacslk/dhcp/dhcpv4/server4"
)

func handler(conn net.PacketConn, peer net.Addr, m *dhcpv4.DHCPv4) {
	// this function will just print the received DHCPv4 message, without replying
	log.Println("handler")
	log.Println(m.Summary())
}

func main() {
	iface := os.Args[1]
	addr := os.Args[2]
	log.Println("iface:", iface)
	log.Println("addr:", addr)

	laddr := net.UDPAddr{
		IP:   net.ParseIP(addr),
		Port: 67,
	}

	server, err := server4.NewServer(iface, &laddr, handler)
	if err != nil {
		log.Fatal(err)
	}

	// This never returns. If you want to do other stuff, dump it into a
	// goroutine.
	server.Serve()
}
