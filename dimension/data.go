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
	Dimension Dimension
	// BlockData returns the datas of the blocks in the world
	BlockData byte
	// PlayerDimensionData returns any data that player has sent while he was in a specific world
	PlayerDimensionData uint64
}
type Generation struct {
	// Dimension...
	Dimension Dimension
	// GenerationDimensionType returns the type of the generation
	GenerationDimensionType uint8
	// Ore returns the ore value
	Ore byte
	// Ore value returns the generation quantity of each ore
	OreValue int32
	// GenerationOreType is the ore that will be eventually modified in OreValue
	GenerationOreType uint8
}
