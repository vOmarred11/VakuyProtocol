package packs

import "github.com/google/uuid"

type packName struct{ string }

type packDescription struct{ string }

type packAuthor struct{ string }

type packVersion struct{ int }

type packFormat struct{ int }

func MainfestName(pc packName) string {
	return pc.string
}
func MainfestDescription(pc packDescription) string {
	return pc.string
}
func MainfestAuthor(pc packAuthor) string {
	return pc.string
}
func MainfestVersion(pc packVersion) int {
	return pc.int
}
func MainfestUUID() uuid.UUID {
	return uuid.UUID{}
}
func MainfestFormat(pc packFormat) int {
	return pc.int
}
