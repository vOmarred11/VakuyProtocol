package pack

type PackName struct{ string }

type PackDescription struct{ string }

type PackVersion struct{ int }

type PackFormat struct{ int }

func (p PackName) PackName() string               { return p.string }
func (p PackDescription) PackDescription() string { return p.string }
func (p PackVersion) PackVersion() int            { return p.int }
func (p PackFormat) PackFormat() int              { return p.int }
