package PhMailStrategy

import (
	"PhSandBox/PhModel"
	"errors"
	"fmt"
)

type MailContext struct {
	strategy MailInterface
	MailModel 	PhModel.Mail
}

func (mc * MailContext) mapping() error {
	var err error
	switch mc.MailModel.Type {
	case 0:
		mc.strategy = &WebUploadEndEmailStrategy{}
	case 1:
		mc.strategy = &ProcessUploadEndEmailStrategy{}
	case 2:
		mc.strategy = &ProcessEndEmailStrategy{}
	case 3:
		mc.strategy = &ErrorEmailStrategy{}
	default:
		err = errors.New(fmt.Sprint("is not implementation"))
	}
	return err
}

func (mc *MailContext) DoExec() (interface{}, error) {
	err := mc.mapping()
	if err != nil {
		return nil, err
	}
	return mc.strategy.DoExec(mc.MailModel)
}
