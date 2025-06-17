package delorian

import (
	"VakuyProtocol/protocol"
	vakuy "VakuyProtocol/server"
	"github.com/pelletier/go-toml"
	"os"
)

var srv = vakuy.Config{}.New()

type ServerConfig struct {
	IP string
}

func server() {
	cfg := ReadServerConfig()
	if cfg.IP == "" {
		cfg.IP = "0.0.0.0:19132"
	}
	srv.Listen()
	for p := range srv.Accept() {
		_ = p
		ServerConnection(protocol.Proto{})
	}
}
func ServerConnection(p protocol.Proto) {
	b, err := p.SendValue(vakuy.Response{})
	if err != nil {
		panic(err)
	}
	err = p.ReceiveValue(b)
	if err != nil {
		panic(err)
	}
	go func() {
		defer func() {
			srv.Close()
		}()
		for player := range srv.Accept() {
			value, err := p.SendValue(player.Tx())
			if err != nil {
				return
			}
			p.ReceiveValue(value)
		}
	}()
}
func ReadServerConfig() ServerConfig {
	b, err := toml.Marshal(ServerConfig{})
	if err != nil {
		panic(err)
	}
	os.WriteFile("server.toml", b, 0777)
	return ServerConfig{}
}
