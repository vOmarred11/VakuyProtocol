package delorian

import (
	"github.com/pelletier/go-toml"
	"github.com/vOmarred11/VakuyProtocol/protocol"
	vakuy "github.com/vOmarred11/VakuyProtocol/server"
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
	defer func() {
		x, err := p.ReadByte()
		if err != nil {
			panic(err)
		}
		err = p.WriteByte(x)
		if err != nil {
			panic(err)
		}
	}()
	go func() {
		defer func() {
			err := srv.Close()
			if err != nil {
				panic(err)
			}
		}()
		for player := range srv.Accept() {
			value, err := p.SendValue(player.Tx())
			if err != nil {
				return
			}
			err = p.ReceiveValue(value)
			if err != nil {
				panic(err)
			}
		}
	}()
}
func ReadServerConfig() ServerConfig {
	b, err := toml.Marshal(ServerConfig{})
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("server.toml", b, 0777)
	if err != nil {
		panic(err)
	}
	return ServerConfig{}
}
