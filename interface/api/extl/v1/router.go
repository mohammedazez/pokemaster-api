package v1

import (
	"log"
	healthCheckService "pokemaster-api/core/healthCheck"
	healtCheckRepositoryMysql "pokemaster-api/infrastructure/repository/mysql/healthCheck"
	healthCheckHandlerV1 "pokemaster-api/interface/api/extl/v1/healthCheck"
	"pokemaster-api/utils/config/mysql"

	"github.com/labstack/echo/v4"
)

func API(route *echo.Echo) {

	// instance mysql
	mysqlDB, err := mysql.GetMysqlDB()
	if err != nil {
		log.Panic(err.Error())
	}

	// instance repo
	healthCheckRepositoryMysql := healtCheckRepositoryMysql.NewRepository(mysqlDB)

	// instance service
	healthCheckService := healthCheckService.NewService(healthCheckRepositoryMysql)

	// instance handler
	healthCheckHandlerV1 := healthCheckHandlerV1.NewHandler(healthCheckService)

	// V1 routes
	v1Route := route.Group("/api/v1")

	healthCheckRouteV1 := v1Route.Group("/health")
	healthCheckRouteV1.GET("", healthCheckHandlerV1.Get)
}
