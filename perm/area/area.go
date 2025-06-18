package area

import (
	"github.com/vOmarred11/VakuyProtocol/player"
	"github.com/vOmarred11/VakuyProtocol/pos"
)

type Area struct {
	StartPos pos.BlockPos
	EndPos   pos.BlockPos
}
type Permission struct {
	Player        *player.Player
	HasPermission bool
}

func (a *Area) AreaStartPos() pos.BlockPos {
	return a.StartPos
}
func (a *Area) AreaEndPos() pos.BlockPos {
	return a.EndPos
}
func (p *Permission) AreaHasPermission() bool {
	ps := Position{}
	pp := player.Player{}
	if p.HasPermission == false {
		// you can handle however you want this
		pp.Message("You've no permissions")
		return false
	}
	if ps.error == nil {
		panic(ps.error.Error())
	}
	return true
}
