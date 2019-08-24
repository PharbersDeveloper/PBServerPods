package State

import (
	"SandBox/Model"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmMongodb"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmRedis"
	"strings"
)

type GroupState struct {
	context *AuthContext
	db *BmMongodb.BmMongodb
	rd *BmRedis.BmRedis
	fileModel Model.FileMetaDatum
}

func (gs *GroupState) DoExecute() (interface{}, error) {
	mod := gs.fileModel.Mod[4:7]
	gs.ChangeState(strings.Contains(mod, gs.context.action))
	return gs.context.state.DoExecute()
}
func (gs *GroupState) ChangeState(status bool) {
	if status { gs.context.state = &EndState{}
	} else { gs.context.state = &NonePermissionsState{} }
}