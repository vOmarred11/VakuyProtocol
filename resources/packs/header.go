package packs

type headerFormatVersion struct{ int }

func FormatVersion(h headerFormatVersion) int {
	return h.int
}
