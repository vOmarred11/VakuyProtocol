package packs

type engineMinVersion struct{ int }

func MinEngineVersion(m engineMinVersion) int {
	return m.int
}
