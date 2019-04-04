package main

import (
	"banwire/dash/dashboard_banwire/net"
//	conf "banwire/dash/dashboard_banwire/conf"
	"flag"
	"log"
	"sync"
)

var loaded sync.WaitGroup

const (
	DefaultRunMode = ""
	ApiRunMode     = "api"
	BatchRunMode   = "batch"
)

func init() {
	loaded.Add(1)
	
}

func main() {
	flag.Parse()
	LoadConfiguration()

	loaded.Done()
	log.Print("Service is ready for run...")

	switch RunMode {
	case BatchRunMode:
		BatchTest()
		break
	case DefaultRunMode, ApiRunMode:
		startServer()
		break
	default:
		log.Print("Run mode is unknown")
		break
	}

}



func startServer() {
	log.Printf("HTTP Listening: %s", HTTPListen)
	net.HTTPListener(HTTPListen)
}
