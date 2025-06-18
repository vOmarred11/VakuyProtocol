package players

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/pelletier/go-toml"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/vOmarred11/VakuyProtocol/server"
)

type Target struct {
	target    byte
	trackList interface {
		targetCamera() mgl32.Vec3
		targetViewPoint() float64
	}
	callbackTrack interface {
		clientCamera() mgl32.Vec3
		clientViewPoint() float64
	}
	stackList struct {
		cameraFromClient mgl32.Vec3
		cameraToClient   mgl32.Vec3
	}
	constraint []byte
}

func (t Target) Data() server.ParsePlayer {
	s := server.ParsePlayer{}
	return s
}
func (t Target) TargetTrackList() []byte {
	b, err := toml.Marshal(t.trackList)
	if err != nil {
		panic(err)
	}
	x := append(b, t.constraint...)
	return x
}
func (t Target) TargetViewPoint() float64 {
	p := t.trackList.targetViewPoint()
	xp, yp, zp := protocol.CameraPreset{}.PosX, protocol.CameraPreset{}.PosY, protocol.CameraPreset{}.PosZ
	_, errx := toml.Marshal(xp)
	_, erry := toml.Marshal(yp)
	_, errz := toml.Marshal(zp)
	if errx != nil || erry != nil || errz != nil {
		panic("error marshalling camera")
	}
	_ = t.constraint
	_ = t.constraint
	_ = t.constraint
	return p
}
