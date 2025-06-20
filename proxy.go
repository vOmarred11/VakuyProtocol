package delorian

import (
	"errors"
	"fmt"
	"github.com/pelletier/go-toml"
	"github.com/vOmarred11/VakuyProtocol/minecraft"
	vakuy "github.com/vOmarred11/VakuyProtocol/protocol"
	"os"
	"time"
)

type ProxyConfig struct {
	Connection struct {
		RemoteAddress string
		LocalAddress  string
	}
}

func proxy() {
	cfg := ReadProxyConfig()
	fmt.Println("Locating external server")
	listener, err := minecraft.ListenConfig{}.Listen("raknet", cfg.Connection.RemoteAddress)
	if err != nil {
		panic(err)
	}
	if errors.Is(err, listener.Close()) {
		panic(errors.New("connection closed"))
	}
	if err != nil {
		panic(err)
	}
	go ProxyConnection(minecraft.Dialer{}, ProxyConfig{}, listener, vakuy.Proto{})

}
func ProxyConnection(dialer minecraft.Dialer, cfg ProxyConfig, listener *minecraft.Listener, p vakuy.Proto) {
	conn, err := dialer.DialTimeout("raknet", cfg.Connection.RemoteAddress, time.Second*10)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to ", cfg.Connection.RemoteAddress)
	go func() {
		if err := conn.StartGame(conn.GameData()); err != nil {
			panic(err)
		}
	}()
	go func() {
		if err := conn.DoSpawn(); err != nil {
			panic(err)
		}
	}()
	go func() {
		defer func() {
			err := listener.Disconnect(conn, "disconnected")
			if err != nil {
				panic(err)
			}
			fmt.Println("disconnected")
		}()
		pk, err := conn.ReadPacket()
		if err != nil {
			panic(err)
		}
		err = conn.WritePacket(pk)
		if err != nil {
			panic(err)
		}
	}()
	go func() {
		defer func() {
			err := listener.Disconnect(conn, "disconnected")
			if err != nil {
				panic(err)
			}
		}()
		pk, err := conn.ReadPacket()
		if err != nil {
			panic(err)
		}
		v, err := p.SendValue(pk)
		if err != nil {
			panic(err)
		}
		err = p.ReceiveValue(v)
		if err != nil {
			panic(err)
		}
	}()
}
func ReadProxyConfig() ProxyConfig {
	b, err := toml.Marshal(ProxyConfig{})
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("proxy.toml", b, 0777)
	if err != nil {
		panic(err)
	}
	return ProxyConfig{}
}
