package main

import (
	"flag"
	"fmt"
	"net"
	"sync"
)

// Get variable from command line
// Usage: go run net/port.go --site=www.url.com
var site = flag.String("site", "scanme.nmap.org", "Site to scan")

func main() {
	flag.Parse() // parse the flags

	var wg sync.WaitGroup

	// Escanear cada puerto y hacer una conexión
	for i := 0; i < 65535; i++ {
		wg.Add(1)
		go func(port int) {
			defer wg.Done()

			conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", *site, port))
			if err != nil {
				return
			}
			conn.Close()
			fmt.Println("Port", port, "is open")
		}(i)
	}

	wg.Wait()
}
