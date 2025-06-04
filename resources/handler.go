package delorian

import (
	"github.com/sandertv/gophertunnel/minecraft/resource"
	"go/types"
)

var resourceType resource.Pack

type Resources struct {
	Name        resource.Pack
	Mainfest    resource.Manifest
	Header      resource.Header
	Dependecies []resource.Dependency
	Capability  resource.Capability
	MetaData    resource.Metadata
	Module      resource.Module
	Modules     []resource.Module
}

func Name(r Resources) (resource.Pack, string) {
	return r.Name, r.Header.Name
}

func Mainfest(r Resources) resource.Manifest {
	return r.Mainfest
}
func Header(r Resources) (resource.Header, string) {
	return r.Header, r.Header.Name
}
func Dependecies(r Resources) []resource.Dependency {
	return r.Dependecies
}
func Capability(r Resources) (resource.Capability, types.Array) {
	return r.Capability, types.Array{}
}
func MetaData(r Resources) (resource.Metadata, Resources) {
	return r.MetaData, Resources{}
}
func Module(r Resources) resource.Module {
	return r.Module
}
func Modules(r Resources) ([]resource.Module, types.Array) {
	return r.Modules, types.Array{}
}
