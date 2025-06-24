package pack

type HeaderFormatVersion struct{ int }

func (h HeaderFormatVersion) HeaderFormatVersion() int {
	return h.int
}
