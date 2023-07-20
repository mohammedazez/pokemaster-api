package v1

import (
	"log"
	healthCheckService "pokemaster-api/core/healthCheck"
	pokemonService "pokemaster-api/core/pokemon"
	userService "pokemaster-api/core/user"
	healthCheckRepositoryMysql "pokemaster-api/infrastructure/repository/mysql/healthCheck"
	pokemonRepositoryMysql "pokemaster-api/infrastructure/repository/mysql/pokemon"
	userRepositoryMysql "pokemaster-api/infrastructure/repository/mysql/user"
	healthCheckHandlerV1 "pokemaster-api/interface/api/extl/v1/healthCheck"
	pokemonHandlerV1 "pokemaster-api/interface/api/extl/v1/pokemon/handler"
	userHandlerV1 "pokemaster-api/interface/api/extl/v1/user/handler"
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
	healthCheckRepositoryMysql := healthCheckRepositoryMysql.NewRepository(mysqlDB)
	pokemonRepositoryMysql := pokemonRepositoryMysql.NewRepository(mysqlDB)
	userRepositoryMysql := userRepositoryMysql.NewRepository(mysqlDB)

	// instance service
	healthCheckService := healthCheckService.NewService(healthCheckRepositoryMysql)
	pokemonService := pokemonService.NewService(pokemonRepositoryMysql)
	userService := userService.NewService(userRepositoryMysql)

	// instance handler
	healthCheckHandlerV1 := healthCheckHandlerV1.NewHandler(healthCheckService)
	pokemonHandlerV1 := pokemonHandlerV1.NewHandler(pokemonService)
	userHandlerV1 := userHandlerV1.NewHandler(userService)

	// V1 routes
	v1Route := route.Group("/api/v1")

	healthCheckRouteV1 := v1Route.Group("/health")
	healthCheckRouteV1.GET("", healthCheckHandlerV1.Get)

	// pokemon api
	pokemonV1 := v1Route.Group("/pokemon")
	pokemonV1.POST("/release-pokemon", pokemonHandlerV1.Insert)
	pokemonV1.GET("/catch-probability", pokemonHandlerV1.CatchPokemon)
	pokemonV1.PUT("/rename-pokemon/:id", pokemonHandlerV1.Update)
	pokemonV1.GET("/list-pokemon", pokemonHandlerV1.List)

	// user api
	userV1 := v1Route.Group("/users")
	userV1.POST("", userHandlerV1.Insert)
	userV1.GET("", userHandlerV1.Get)
}
