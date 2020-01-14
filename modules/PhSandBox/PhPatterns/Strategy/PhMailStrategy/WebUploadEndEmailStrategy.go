package PhMailStrategy

import "PhSandBox/PhModel"

type WebUploadEndEmailStrategy struct {}

func (w * WebUploadEndEmailStrategy) DoExec(mail PhModel.Mail) (interface{}, error) {
	return nil, nil
}
