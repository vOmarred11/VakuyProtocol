package server

import (
	"VakuyServer/player"
	"github.com/pelletier/go-toml"
)

type ParsePlayer struct {
	incoming chan interface {
		Data() byte
		ParsedData() []byte
	}
	outgoing chan interface {
		Data() byte
		ParsedData() []byte
		Changes() chan bool
	}
	callbackStack interface {
		GameStack() byte
		CallBackType() uint8
		Signal() map[uint8]byte
	}
	callbackTrace interface {
		GameTrace() byte
		CallbackStack() uint8
		Signal() map[uint8]byte
	}

	client map[interface {
		Data() chan byte
		ParsedData() []byte
		ServerParsedData() []uintptr
	}]interface {
		Outgoing() chan interface {
			Data() chan byte
		}
	}
	players map[struct {
		sig byte
	}]byte
	ph func(outdata byte)
}
type Response struct {
	responseMovement, responseCamera, responseInventory, responseHit byte
}

func (r ParsePlayer) ResponseMovement() ([]byte, []byte) {
	b, err := toml.Marshal(r.callbackTrace.CallbackStack())
	if err != nil {
		panic(err)
	}
	s := toml.Decoder{}
	err = s.Decode(r.ph)
	if err != nil {
		panic(err)
	}
	x := append([]byte{r.callbackTrace.CallbackStack()}, b...)
	e, err := toml.Marshal(s)
	if err != nil {
		panic(err)
	}
	if x == nil {
		panic(err)
	}
	return x, e
}
func (r ParsePlayer) ResponseCamera() ([]byte, []byte) {
	b, err := toml.Marshal(r.callbackTrace.CallbackStack())
	if err != nil {
		panic(err)
	}
	s := toml.Decoder{}
	err = s.Decode(r.ph)
	if err != nil {
		panic(err)
	}
	x := append([]byte{r.callbackStack.GameStack()}, b...)
	e, err := toml.Marshal(s)
	if err != nil {
		panic(err)
	}
	if x == nil {
		panic(err)
	}
	return x, e
}
func (r ParsePlayer) ResponseHit() ([]byte, []byte) {
	b, err := toml.Marshal(r.callbackTrace.CallbackStack())
	if err != nil {
		panic(err)
	}
	s := toml.Decoder{}
	err = s.Decode(r.ph)
	if err != nil {
		panic(err)
	}
	x := append([]byte{r.callbackTrace.CallbackStack()}, b...)
	e, err := toml.Marshal(s)
	if err != nil {
		panic(err)
	}
	if x == nil {
		panic(err)
	}
	return x, e
}
func (r ParsePlayer) ResponseInventory() ([]byte, []byte) {
	b, err := toml.Marshal(r.callbackTrace.CallbackStack())
	if err != nil {
		panic(err)
	}
	s := toml.Decoder{}
	err = s.Decode(r.ph)
	if err != nil {
		panic(err)
	}
	x := append([]byte{r.callbackStack.CallBackType()}, b...)
	e, err := toml.Marshal(s)
	if err != nil {
		panic(err)
	}
	if x == nil {
		panic(err)
	}
	return x, e
}
func (r ParsePlayer) ResponseTarget() player.Config {
	b, err := toml.Marshal(r.ph)
	if err != nil {
		panic(err)
	}
	s := toml.Decoder{}
	s.Decode(b)
	p := player.Player{}
	return p.Data()
}
