package PhMailStrategy

import "PhSandBox/PhModel"

type ProcessUploadEndEmailStrategy struct {}

func (p *ProcessUploadEndEmailStrategy) DoExec(mail PhModel.Mail) (interface{}, error) {
	return nil, nil
}