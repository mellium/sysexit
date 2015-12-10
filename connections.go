package honey

import (
	"net"

	"bitbucket.org/SamWhited/honey/xmpp"
)

// C2S contains configuration for client to server connections.
type VirtualHost struct {
	Domain string  `toml:"domain"`
	C2S    Service `toml:"c2s"`
}

// A Service is anything for which we should spawn a listener.
type Service struct {
	Addr     string `toml:"addr"`
	Certfile string `toml:"certfile"`
	Keyfile  string `toml:"keyfile"`
}

// VHostsListenAndServe starts listening for configured C2S and S2S connections.
func VHostsListenAndServe(vhosts []VirtualHost) error {
	for _, v := range vhosts {
		addr := v.C2S.Addr
		server := &xmpp.Server{
			ClientAddr: addr,
			Handler:    nil, // TODO
			TLSConfig:  nil, // TODO
		}

		listener, err := net.Listen("tcp", addr)
		if err != nil {
			return err
		}
		defer func() {
			if cerr := listener.Close(); cerr != nil {
				err = cerr
			}
		}()
		go server.Serve(listener)
		return err
	}

	select {}
}
