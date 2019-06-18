package main

import (
	"github.com/micro/go-micro"

	"github.com/fidelfly/fxms/mspkg/db"
	"github.com/fidelfly/fxms/srv/user/config"
	"github.com/fidelfly/fxms/srv/user/pwd"
	"github.com/fidelfly/fxms/srv/user/res"
)

func Initializer() micro.Option {
	return func(options *micro.Options) {
		err := initSa(config.Sa)
		if err != nil {
			panic("Initial Sa failed")
		}
	}
}

func initSa(saCfg *config.SuperUser) error {
	sa := res.User{}
	if len(saCfg.Code) > 0 {
		sa.Code = saCfg.Code
	} else {
		sa.Code = "sa"
	}
	if len(saCfg.Email) > 0 {
		sa.Email = saCfg.Email
	}

	if find, err := db.Read(&sa); err != nil {
		return err
	} else if find {
		return nil
	} else {
		sa.Name = saCfg.Name
		sa.Password = pwd.EncodePwd(sa.Code, saCfg.Password)
		_, err = db.Create(sa)
		return err
	}
}
