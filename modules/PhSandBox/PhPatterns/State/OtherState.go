package State

import (
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmMongodb"
)

type OtherState struct {
	context *AuthContext
	db *BmMongodb.BmMongodb
}

func (os *OtherState) DoExecute() (interface{}, error) {
	//if len(os.assetsModel.Accessibility) == 0 { os.ChangeState(false); return os.context.state.DoExecute() }
	//mod := os.assetsModel.Accessibility[7:]
	//os.ChangeState(strings.Contains(mod, os.context.action))
	return os.context.state.DoExecute()
}
func (os *OtherState) ChangeState(status bool) {
	if status { os.context.state = &EndState{}
	} else { os.context.state = &NonePermissionsState{} }
}