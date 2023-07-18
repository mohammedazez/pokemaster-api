package handler

import (
	"net/http"
	port "pokemaster-api/core/port/pokemon"
	"pokemaster-api/interface/api/extl/v1/pokemon/request"
	"pokemaster-api/interface/api/extl/v1/pokemon/response"
	"pokemaster-api/interface/api/utils/validation"

	domain "pokemaster-api/core/domain/pokemon"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	service port.Service
}

func NewHandler(service port.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Insert(c echo.Context) error {

	req := new(request.RequestInsert)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	errVal := validation.ValidateReq(req)
	if errVal != nil {
		return c.JSON(http.StatusBadRequest, errVal)
	}

	pokemon := new(domain.Pokemon)
	pokemon.PokemonName = req.PokemonName
	pokemon.PokemonPicture = req.PokemonPicture
	pokemon.Number = req.Number

	result, err := h.service.Insert(pokemon)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	rest := new(response.PrimeNumberCheck)
	if !isPrime(result.Number) {
		rest.Result = "failure"
		rest.Released = false
		rest.PrimeNumber = result.Number
	} else {
		rest.Result = "success"
		rest.Released = true
		rest.PrimeNumber = result.Number
	}

	res := response.NewResponseDetail(http.StatusText(http.StatusCreated), rest, http.StatusCreated, true)
	return c.JSON(http.StatusCreated, res)
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}
