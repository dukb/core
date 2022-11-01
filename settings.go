package core

import (
	"github.com/dukb/core/config"
	"log"
)

type Settings struct {
	Settings  config.Config `yaml:"settings"`
	callbacks []func()
}

func (e *Settings) runCallback() {
	for i := range e.callbacks {
		e.callbacks[i]()
	}
}

func (e *Settings) OnChange() {
	e.init()
	log.Println("!!! config change and reload")
}

func (e *Settings) Init() {
	e.init()
	log.Println("!!! config init")
}

func (e *Settings) init() {
	e.Settings.Logger.Setup()
	e.runCallback()
}
