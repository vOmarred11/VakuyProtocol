package dimension

const (
	GenerationData = iota
	GenerationOreCoal
	GenerationOreIron
	GenerationOreGold
	GenerationOreDiamond
	GenerationBiome
	GenerationTypeNormal
	GenerationTypeAmplified
	GenerationTypeCustom
	GenerationTypeFlat
)

type Data struct {
	// Dimension...
	DimensionData Dimension
	// BlockData returns the datas of the blocks in the world
	BlockData byte
	// PlayerDimensionData returns any data that player has sent while he was in a specific world
	PlayerDimensionData uint64
}
type Generation struct {
	// Dimension...
	Dimension Dimension
	// DimensionType returns the type of the generation
	DimensionType uint8
	// Ore returns the ore value
	Ore byte
	// Ore value returns the generation quantity of each ore
	OreValue int32
	// OreType is the ore that will be eventually modified in OreValue
	OreType uint8
}

func (d Data) DataDimensionData() Dimension {
	return d.DimensionData
}
func (d Data) DataBlockData() byte {
	return d.BlockData
}
func (d Data) DataPlayerDimensionData() uint64 {
	return d.PlayerDimensionData
}
func (g Generation) GenerationDimension() Dimension {
	return g.Dimension
}
func (g Generation) GenerationDimensionType() uint8 {
	return g.DimensionType
}
func (g Generation) GenerationOre() byte {
	return g.Ore
}
func (g Generation) GenerationOreValue() int32 {
	return g.OreValue
}
func (g Generation) GenerationOreType() uint8 {
	return g.OreType
}
