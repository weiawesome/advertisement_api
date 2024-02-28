package main

import (
	"advertisement_api/api/routes"
	"advertisement_api/utils"
)

func main() {
	if err := utils.InitDB(); err != nil {
		//utils.LogFatal(utils.LogData{Event: "Failed to connect to the database", User: "system", Details: err.Error()})
		return
	}
	defer func() {
		err := utils.CloseDB()
		if err != nil {
			//utils.LogFatal(utils.LogData{Event: "Failed to disconnect to the database", User: "system", Details: err.Error()})
		}
	}()
	if err := utils.InitRedis(); err != nil {
		//utils.LogFatal(utils.LogData{Event: "Failed to connect to the redis", User: "system", Details: err.Error()})
		return
	}
	defer func() {
		err := utils.CloseRedis()
		if err != nil {
			//utils.LogFatal(utils.LogData{Event: "Failed to disconnect to the redis", User: "system", Details: err.Error()})
		}
	}()

	utils.InitMaps()
	utils.InitSingleFLight()

	r := routes.InitRoutes()
	err := r.Run()

	if err != nil {
		return
	}
}
