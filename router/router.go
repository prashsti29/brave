package router

import (
	"github.com/gorilla/mux"
	"github.com/prashsti29/brave/controllers"
)

func SetupRouter(
	playerController *controllers.PlayerController,
	buildingController *controllers.BuildingController,
	villageController *controllers.VillageLayoutController,
) *mux.Router {

	var appRouter *mux.Router
	appRouter = mux.NewRouter()

	RegisterPlayerRoutes(appRouter, playerController)
	RegisterBuildingRoutes(appRouter, buildingController)
	RegisterVillageRoutes(appRouter, villageController)

	return appRouter
}
