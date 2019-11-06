package State

type IdentifyState struct {
	context *AuthContext
}

func (os *IdentifyState) DoExecute() (interface{}, error) {
	if len(os.context.realMod) == 0 {
		os.ChangeState(false)
		return os.context.state.DoExecute()
	}

	// TODO ： 这边差个判断
	os.ChangeState(os.context.action == os.context.realMod)
	return os.context.state.DoExecute()

	//mod := os.assetsModel.Accessibility[1:4]
	//os.ChangeState(strings.Contains(mod, os.context.action))
	//return os.context.state.DoExecute()
}

func (os *IdentifyState) ChangeState(status bool) {
	if status { os.context.state = &EndState{}
	} else { os.context.state = &NonePermissionsState{} }
}