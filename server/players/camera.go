package players

import "github.com/go-gl/mathgl/mgl32"

const (
	FirstPerson = iota
	SecondPerson
	ThirdPerson
)

// Camera is the player camera view
type Camera struct {
	// Angle is the 360° angle of the camera
	Angle mgl32.Vec3
	// Person defines what person is player on
	Person uint8
	// Where defines where the player is looking at
	Where float32
	// LookingAtPlayer defines if the player is looking at the client
	LookingAtPlayer bool
}

func (c *Camera) CameraAngle() mgl32.Vec3 {
	return c.Angle
}
func (c *Camera) CameraPerson() uint8 {
	return c.Person
}
func (c *Camera) CameraWhere() float32 {
	return c.Where
}
func (c *Camera) CameraLookingAtPlayer() bool {
	return c.LookingAtPlayer
}
