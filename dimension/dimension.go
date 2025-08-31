package dimension

type Dimension struct {
	Data byte
}

func (d Dimension) DimensionData() byte {
	return d.Data
}
