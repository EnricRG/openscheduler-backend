package app

import "github.com/EnricRG/openscheduler-backend/internal/service"

type Application struct {
	Services service.Services
}

func (a *Application) Start() {

}

func (a *Application) Stop() {

}
