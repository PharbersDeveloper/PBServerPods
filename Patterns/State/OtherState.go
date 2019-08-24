package State

import (
	"SandBox/Model"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmMongodb"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmRedis"
	"strings"
)

type OtherState struct {
	context *AuthContext
	db *BmMongodb.BmMongodb
	rd *BmRedis.BmRedis
	fileModel Model.FileMetaDatum
}

func (os *OtherState) DoExecute() (interface{}, error) {
	mod := os.fileModel.Mod[7:]
	os.ChangeState(strings.Contains(mod, os.context.action))
	return os.context.state.DoExecute()
}
func (os *OtherState) ChangeState(status bool) {
	if status { os.context.state = &EndState{}
	} else { os.context.state = &NonePermissionsState{} }
}