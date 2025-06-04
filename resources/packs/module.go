package packs

import "github.com/google/uuid"

type moduleDescription struct{ string }

type moduleVersion struct{ int }

type moduleUUID struct{ uuid.UUID }

type moduleType struct{ string }

func getModuleDescription(md moduleDescription) string {
	return md.string
}
func getModuleVersion(md moduleVersion) int {
	return md.int
}
func getModuleUUID(md moduleUUID) uuid.UUID {
	return md.UUID
}
func getType(md moduleType) string {
	md.string = "resources"
	return md.string
}
