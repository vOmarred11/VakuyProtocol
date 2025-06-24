package pack

import "github.com/google/uuid"

type ModuleDescription struct{ string }

type ModuleVersion struct{ int }

type ModuleUUID struct{ uuid.UUID }

type ModuleType struct{ string }

func (m ModuleDescription) ModuleDescription() string { return m.string }
func (m ModuleVersion) ModuleVersion() int            { return m.int }
func (m ModuleUUID) ModuleUUID() uuid.UUID            { return m.UUID }
func (m ModuleType) ModuleType() string {
	m.string = "resources"
	return m.string
}
