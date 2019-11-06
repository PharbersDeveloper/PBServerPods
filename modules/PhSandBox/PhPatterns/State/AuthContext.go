package State

type AuthContext struct {
	state State
	target string
	who map[string]string
	action string
	realMod string
}

func (c *AuthContext) NewAuthContext(target,
									action string,
									who map[string]string,
									realMod string) {
	c.target = target
	c.who = who
	c.action = action
	c.realMod = realMod
}

func (c *AuthContext) mappingState() {
	c.state = &IdentifyState{c}

	//if c.who["account"] == out.Owner {
	//	c.state = &OwnerState{c, c.db, out}
	//} else if c.who["group"] == "" { //out.GroupID {
	//	c.state = &GroupState{c, c.db,out}
	//} else {
	//	c.state = &OtherState{c, c.db,out}
	//}
}

func (c *AuthContext) DoExecute() (interface{}, error) {
	c.mappingState()
	return c.state.DoExecute()
}
