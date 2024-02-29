/*
The main function to start the application.
At first, it initializes all setting include logger, maps, singleflight, database.
Then it starts the service by the framework, gin-gonic/gin.
*/
package main

import (
	"advertisement_api/api/routes"
	"advertisement_api/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	// initialize the logger.
	utils.InitLogger()

	// initialize the maps include gender-map, country-map and platform-map.
	utils.InitMaps()

	// initialize the singleflight instance
	utils.InitSingleFLight()

	// initialize the sql-db connection. return and log error when connect fail.
	if err := utils.InitDB(); err != nil {
		utils.LogFatal(err.Error())
		return
	}
	// close sql-db connection when application turn off. log error when close fail.
	defer func() {
		err := utils.CloseDB()
		if err != nil {
			utils.LogFatal(err.Error())
		}
	}()

	// initialize the redis connection. return and log error when connect fail.
	if err := utils.InitRedis(); err != nil {
		utils.LogFatal(err.Error())
		return
	}
	// close redis connection when application turn off. log error when close fail.
	defer func() {
		err := utils.CloseRedis()
		if err != nil {
			utils.LogFatal(err.Error())
		}
	}()

	// set mode to the release mode.
	gin.SetMode(gin.ReleaseMode)

	// initialize all the routes in the server.
	r := routes.InitRoutes()
	// start the service. log error when start fail.
	err := r.Run()
	if err != nil {
		utils.LogFatal(err.Error())
		return
	}
}
