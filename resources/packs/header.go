package packs

type headerFormatVersion struct{ int }

func getFormatVersion(h headerFormatVersion) int {
	return h.int
}
