package PhMailStrategy

import "PhSandBox/PhModel"

type ErrorEmailStrategy struct {}

func (e *ErrorEmailStrategy) DoExec(mail PhModel.Mail) (interface{}, error) {
	return nil, nil
}