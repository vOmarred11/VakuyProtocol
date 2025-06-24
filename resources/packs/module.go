package packs

import "github.com/google/uuid"

type moduleDescription struct{ string }

type moduleVersion struct{ int }

type moduleUUID struct{ uuid.UUID }

type moduleType struct{ string }

func ModuleDescription(md moduleDescription) string {
	return md.string
}
func ModuleVersion(md moduleVersion) int {
	return md.int
}
func ModuleUUID(md moduleUUID) uuid.UUID {
	return md.UUID
}
func Type(md moduleType) string {
	md.string = "resources"
	return md.string
}
