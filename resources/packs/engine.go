package packs

type EngineMinVersion struct{ int }

func (e EngineMinVersion) EngineMinVersion() int { return e.int }
