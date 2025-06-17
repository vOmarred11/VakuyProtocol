package ohc

// Ohc is any action or else sent by the server
type Ohc struct {
	// Data is the Ohc data
	Data byte
	// Packets is a slice of all packets sent to you
	Packets []byte
	// ToEntities defines if the Ohc is sent also to entities such as Rabbit, Cow...
	ToEntities bool
	// ToEveryone defines if the Ohc is sent to the whole server
	ToEveryone bool
}

func (o Ohc) OhcData() byte {
	return o.Data
}
func (o Ohc) OhcPackets() []byte {
	return o.Packets
}
func (o Ohc) OhcToEntities() bool {
	return o.ToEntities
}
func (o Ohc) OhcToEveryone() bool {
	return o.ToEveryone
}
