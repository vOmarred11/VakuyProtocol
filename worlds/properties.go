package worlds

import "github.com/sirupsen/logrus"

type properties struct {
	BreakBlocks      bool
	PlaceBlocks      bool
	HitPlayers       bool
	UseItems         bool
	InteractEntities bool
	PickItems        bool
	PickupItems      bool
	DropItems        bool
	Fly              bool
	Teleport         bool
	Move             bool
	ExecuteCommands  bool
}

type errl struct{ error }
type ServerResponse struct {
	Default bool
	Data    map[string]interface{}
	Error   error
}

var errlInstance *errl

func BreakBlocks(value bool) ServerResponse {
	if errlInstance == nil {
		logrus.Info("loc: minecraft/delorian/worlds, level: err, err: " + errlInstance.Error())
		return ServerResponse{Error: errlInstance.error}
	}
	return ServerResponse{Default: true, Data: map[string]interface{}{"BreakBlocks": value}}
}
func PlaceBlocks(value bool) ServerResponse {
	if errlInstance == nil {
		logrus.Info("loc: minecraft/delorian/worlds, level: err, err: " + errlInstance.Error())
		return ServerResponse{Error: errlInstance.error}
	}
	return ServerResponse{Default: true, Data: map[string]interface{}{"PlaceBlocks": value}}
}

func HitPlayers(value bool) ServerResponse {
	if errlInstance == nil {
		logrus.Info("loc: minecraft/delorian/worlds, level: err, err: " + errlInstance.Error())
		return ServerResponse{Error: errlInstance.error}
	}
	return ServerResponse{Default: true, Data: map[string]interface{}{"HitPlayers": value}}
}

func UseItems(value bool) ServerResponse {
	if errlInstance == nil {
		logrus.Info("loc: minecraft/delorian/worlds, level: err, err: " + errlInstance.Error())
		return ServerResponse{Error: errlInstance.error}
	}
	return ServerResponse{Default: true, Data: map[string]interface{}{"UseItems": value}}
}

func InteractEntities(value bool) ServerResponse {
	if errlInstance == nil {
		logrus.Info("loc: minecraft/delorian/worlds, level: err, err: " + errlInstance.Error())
		return ServerResponse{Error: errlInstance.error}
	}
	return ServerResponse{Default: true, Data: map[string]interface{}{"InteractEntities": value}}
}

func PickItems(value bool) ServerResponse {
	if errlInstance == nil {
		logrus.Info("loc: minecraft/delorian/worlds, level: err, err: " + errlInstance.Error())
		return ServerResponse{Error: errlInstance.error}
	}
	return ServerResponse{Default: true, Data: map[string]interface{}{"PickItems": value}}
}

func PickupItems(value bool) ServerResponse {
	if errlInstance == nil {
		logrus.Info("loc: minecraft/delorian/worlds, level: err, err: " + errlInstance.Error())
		return ServerResponse{Error: errlInstance.error}
	}
	return ServerResponse{Default: true, Data: map[string]interface{}{"PickupItems": value}}
}

func DropItems(value bool) ServerResponse {
	if errlInstance == nil {
		logrus.Info("loc: minecraft/delorian/worlds, level: err, err: " + errlInstance.Error())
		return ServerResponse{Error: errlInstance.error}
	}
	return ServerResponse{Default: true, Data: map[string]interface{}{"DropItems": value}}
}

func Fly(value bool) ServerResponse {
	if errlInstance == nil {
		logrus.Info("loc: minecraft/delorian/worlds, level: err, err: " + errlInstance.Error())
		return ServerResponse{Error: errlInstance.error}
	}
	return ServerResponse{Default: false, Data: map[string]interface{}{"Fly": value}}
}

func Teleport(value bool) ServerResponse {
	if errlInstance == nil {
		logrus.Info("loc: minecraft/delorian/worlds, level: err, err: " + errlInstance.Error())
		return ServerResponse{Error: errlInstance.error}
	}
	return ServerResponse{Default: true, Data: map[string]interface{}{"Teleport": value}}
}

func Move(value bool) ServerResponse {
	if errlInstance == nil {
		logrus.Info("loc: minecraft/delorian/worlds, level: err, err: " + errlInstance.Error())
		return ServerResponse{Error: errlInstance.error}
	}
	return ServerResponse{Default: true, Data: map[string]interface{}{"Move": value}}
}

func ExecuteCommands(value bool) ServerResponse {
	if errlInstance == nil {
		logrus.Info("loc: minecraft/delorian/worlds, level: err, err: " + errlInstance.Error())
		return ServerResponse{Error: errlInstance.error}
	}
	return ServerResponse{Default: true, Data: map[string]interface{}{"ExecuteCommands": value}}
}
