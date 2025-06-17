package packs

type engineMinVersion struct{ int }

func getMinEngineVersion(m engineMinVersion) int {
	return m.int
}
