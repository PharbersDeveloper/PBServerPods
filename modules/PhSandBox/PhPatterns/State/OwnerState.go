package State

import (
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmMongodb"
)

type OwnerState struct {
	context *AuthContext
	db *BmMongodb.BmMongodb
}

func (os *OwnerState) DoExecute() (interface{}, error) {
	//if len(os.assetsModel.Accessibility) == 0 { os.ChangeState(false); return os.context.state.DoExecute() }
	//
	//mod := os.assetsModel.Accessibility[1:4]
	//os.ChangeState(strings.Contains(mod, os.context.action))
	return os.context.state.DoExecute()
}

func (os *OwnerState) ChangeState(status bool) {
	if status { os.context.state = &EndState{}
	} else { os.context.state = &NonePermissionsState{} }
}