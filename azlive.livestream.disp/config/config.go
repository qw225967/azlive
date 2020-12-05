package config

import (
	"github.com/BurntSushi/toml"
	"github.com/pkg/errors"
)

func Initconf() (FileConf ,error) {
	var conf FileConf
	if _, err := toml.DecodeFile("./conf/config.toml", &conf); err != nil {
		return conf,errors.New("Open File is Failed!")
	}
	return conf,nil
}

