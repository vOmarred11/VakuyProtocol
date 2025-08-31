package dimension

// Dimension returns the dimension data
type Dimension struct {
	Data byte
}

func (d Dimension) DimensionData() byte {
	return d.Data
}
