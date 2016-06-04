package mel

import (
	"crypto/tls"

	"mellium.im/xmpp/server"
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

// ServerFromConfig creates a new Server from the given VirtualHost config.
func ServerFromConfig(vhost VirtualHost) *server.Server {
	cert, err := tls.LoadX509KeyPair(vhost.C2S.Certfile, vhost.C2S.Keyfile)
	if err != nil {
		log.Fatal("Error loading X509 key pair.", err)
	}
	return server.New(
		server.ClientAddr(vhost.C2S.Addr),
		server.TLSConfig(&tls.Config{
			Certificates: []tls.Certificate{cert},
		}),
	)
}
