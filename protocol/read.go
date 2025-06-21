package protocol

func (p *Proto) String() string {
	return string(p.Bytes)
}
func (p *Proto) Byte() byte {
	return p.Bytes[1]
}
func (p *Proto) LenByte() int { return int(p.Bytes[1]) }
func (p *Proto) Int() int {
	return int(p.Amount)
}
func (p *Proto) Int8() int8 {
	return int8(p.Amount)
}
func (p *Proto) Int16() int16 {
	return int16(p.Amount)
}
func (p *Proto) Int32() int32 {
	return p.Amount
}
func (p *Proto) Int64() int64 {
	return int64(p.Amount)
}
func (p *Proto) Uint() uint {
	return uint(p.Amount)
}
func (p *Proto) Uint8() uint8 {
	return uint8(p.Amount)
}
func (p *Proto) Uint16() uint16 {
	return uint16(p.Amount)
}
func (p *Proto) Uint32() uint32 {
	return uint32(p.Amount)
}
func (p *Proto) Uint64() uint64 {
	return uint64(p.Amount)
}
func (p *Proto) Float32() float32 {
	return float32(p.Amount)
}
func (p *Proto) Float64() float64 {
	return float64(p.Amount)
}
