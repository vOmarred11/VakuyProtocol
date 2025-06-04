package packs

import "github.com/google/uuid"

type packName struct{ string }

type packDescription struct{ string }

type packAuthor struct{ string }

type packVersion struct{ int }

type packFormat struct{ int }

func getMainfestName(pc packName) string {
	return pc.string
}
func getMainfestDescription(pc packDescription) string {
	return pc.string
}
func getMainfestAuthor(pc packAuthor) string {
	return pc.string
}
func getMainfestVersion(pc packVersion) int {
	return pc.int
}
func getMainfestUUID() uuid.UUID {
	return uuid.UUID{}
}
func getMainfestFormat(pc packFormat) int {
	return pc.int
}
