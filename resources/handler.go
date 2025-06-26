package delorian

import (
	"github.com/sandertv/gophertunnel/minecraft/resource"
	"go/types"
)

type ResourcePack struct {
	Name        resource.Pack
	Mainfest    resource.Manifest
	Header      resource.Header
	Dependecies []resource.Dependency
	Capability  resource.Capability
	MetaData    resource.Metadata
	Module      resource.Module
	Modules     []resource.Module
}

func Name(r ResourcePack) (resource.Pack, string) {
	return r.Name, r.Header.Name
}

func Mainfest(r ResourcePack) resource.Manifest {
	return r.Mainfest
}
func Header(r ResourcePack) (resource.Header, string) {
	return r.Header, r.Header.Name
}
func Dependecies(r ResourcePack) []resource.Dependency {
	return r.Dependecies
}
func Capability(r ResourcePack) (resource.Capability, types.Array) {
	return r.Capability, types.Array{}
}
func MetaData(r ResourcePack) (resource.Metadata, ResourcePack) {
	return r.MetaData, ResourcePack{}
}
func Module(r ResourcePack) resource.Module {
	return r.Module
}
func Modules(r ResourcePack) ([]resource.Module, types.Array) {
	return r.Modules, types.Array{}
}
