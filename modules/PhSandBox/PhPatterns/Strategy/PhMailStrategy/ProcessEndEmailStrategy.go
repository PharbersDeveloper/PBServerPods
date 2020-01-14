package PhMailStrategy

import "PhSandBox/PhModel"

type ProcessEndEmailStrategy struct {}

func (p *ProcessEndEmailStrategy) DoExec(mail PhModel.Mail) (interface{}, error) {
	return nil, nil
}
