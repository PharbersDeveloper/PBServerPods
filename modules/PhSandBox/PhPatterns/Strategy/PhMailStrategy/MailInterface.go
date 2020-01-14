package PhMailStrategy

import "PhSandBox/PhModel"

type MailInterface interface {
	DoExec(mail PhModel.Mail) (interface{}, error)
}
