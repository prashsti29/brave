package main

import (
"fmt"
"log"
"net/http"

"github.com/prashsti29/brave/config"
"github.com/prashsti29/brave/controllers"
"github.com/prashsti29/brave/repository"
"github.com/prashsti29/brave/router"
"github.com/prashsti29/brave/service"
)

func main() {
	db := config.ConnectDB()
	defer db.Close()

	fmt.Println("Server starting...")

	playerRepo := repository.NewPlayerRepository(db)
	buildingRepo := repository.NewBuildingRepository(db)
	villageRepo := repository.NewVillageLayoutRepository(db)

	var villageService = service.NewVillageLayoutService(villageRepo)

	var buildingService = service.NewBuildingService(buildingRepo, villageService)

	var playerService = service.NewPlayerService(playerRepo, buildingService)

	var playerController = controllers.NewPlayerController(playerService)
	var buildingController = controllers.NewBuildingController(buildingService)
	var villageController = controllers.NewVillageLayoutController(villageService)

	var configService = service.NewConfigService(
repository.NewTroopConfigRepository(db),
repository.NewProducerBuildingConfigRepository(db),
repository.NewStorageBuildingConfigRepository(db),
repository.NewDefenseBuildingConfigRepository(db),
repository.NewLaboratoryConfigRepository(db),
repository.NewArmyCampConfigRepository(db),
repository.NewBarrackConfigRepository(db),
repository.NewBuildingConfigRepository(db),
)

	var configController = controllers.NewConfigController(configService)

	var appRouter = router.SetupRouter(
playerController,
buildingController,
villageController,
configController,
)

	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", appRouter))
}
