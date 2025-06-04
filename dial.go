package delorian

import (
	vakuy "VakuyProtocol/protocol"
	"fmt"
	"github.com/pelletier/go-toml"
	"github.com/sandertv/gophertunnel/minecraft"
	"os"
	"time"
)

type Config struct {
	Connection struct {
		RemoteAddress string
		LocalAddress  string
	}
}

func main() {
	cfg := ReadConfig()
	fmt.Println("Locating external server")
	listener, err := minecraft.ListenConfig{}.Listen("raknet", cfg.Connection.RemoteAddress)
	if err != nil {
		panic(err)
	}
	if listener.Close(); err != nil {
		panic(err)
	}
	go Connection(minecraft.Dialer{}, Config{}, listener, vakuy.Proto{})

}
func Connection(dialer minecraft.Dialer, cfg Config, listener *minecraft.Listener, p vakuy.Proto) {
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
			listener.Disconnect(conn, "disconnected")
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
			listener.Disconnect(conn, "disconnected")
		}()
		pk, err := conn.ReadPacket()
		if err != nil {
			panic(err)
		}
		v := p.SendValue(pk)
		err = p.ReceiveValue(v)
		if err != nil {
			panic(err)
		}
	}()
}
func ReadConfig() Config {
	b, err := toml.Marshal(Config{})
	if err != nil {
		panic(err)
	}
	os.WriteFile("delorian.toml", b, 0777)
	return Config{}
}
