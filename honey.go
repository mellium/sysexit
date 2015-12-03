package main

import (
	"net"

	"./config"
	"./logging"
	"./xmpp-core"
)

func main() {

	c := new(config)

	// Load the config
	err := config.ReloadConfig()
	if err != nil {
		logger.Err(err.Error())
		return
	}

	// Setup logging

	// Spawn S2S listeners
	for _, v := range config.C.S2S.Services {
		addr, err := net.ResolveTCPAddr("tcp", v.Addr)
		if err != nil {
			logger.Err(err.Error())
			return
		}
		listener, err := net.ListenTCP("tcp", addr)
		if err != nil {
			logger.Err(err.Error())
			return
		}
		defer func() {
			if cerr := listener.Close(); err == nil {
				err = cerr
			}
		}()
		go listen(listener)
	}

	// Spawn C2S listeners
	for _, v := range config.C.C2S.Services {
		addr, err := net.ResolveTCPAddr("tcp", v.Addr)
		if err != nil {
			logger.Err(err.Error())
			return
		}
		listener, err := net.ListenTCP("tcp", addr)
		if err != nil {
			logger.Err(err.Error())
			return
		}
		defer func() {
			if cerr := listener.Close(); err == nil {
				err = cerr
			}
		}()
		go listen(listener)
	}

	// Block forever.
	select {}
}

// Listen for connection attempts and spawn handlers to deal with them.
func listen(l net.Listener) error {
	// TODO: Don't allow infinite connections.
	for {
		conn, err := l.Accept()
		if err != nil {
			logger.Debug(err.Error())
		} else {
			// Serve connections concurrently.
			go xmpp.Handle(conn, l)
		}
	}
}
