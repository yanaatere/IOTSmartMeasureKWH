package main

import (
	repositories "github.com/retere/IOTSmartMeasureKWH/model"
)

func (se *Serve) initializeRoutes() {
	repository := repositories.NewRepository(se.DB)

	tenant := tenant.NewSer

}
