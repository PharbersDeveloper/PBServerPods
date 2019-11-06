package State

import (
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmMongodb"
)

type GroupState struct {
	context *AuthContext
	db *BmMongodb.BmMongodb
}

func (gs *GroupState) DoExecute() (interface{}, error) {
	//if len(gs.assetsModel.Accessibility) == 0 { gs.ChangeState(false); return gs.context.state.DoExecute() }
	//mod := gs.assetsModel.Accessibility[4:7]
	//gs.ChangeState(strings.Contains(mod, gs.context.action))
	return gs.context.state.DoExecute()
}
func (gs *GroupState) ChangeState(status bool) {
	if status { gs.context.state = &EndState{}
	} else { gs.context.state = &NonePermissionsState{} }
}