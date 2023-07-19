package handler

import (
	"net/http"
	port "pokemaster-api/core/port/pokemon"
	"pokemaster-api/interface/api/common"
	"pokemaster-api/interface/api/extl/v1/pokemon/request"
	"pokemaster-api/interface/api/extl/v1/pokemon/response"
	"pokemaster-api/interface/api/utils/validation"
	"strings"

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
		if strings.Contains(err.Error(), "number is not prime") {
			return echo.NewHTTPError(http.StatusBadRequest, "number must be a prime number")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	resp := new(response.PrimeNumberCheck)
	resp.Result = "success"
	resp.Released = true
	resp.PrimeNumber = result.Number

	res := new(common.DefaultResponse)
	res.SetResponseData(http.StatusText(http.StatusOK), resp, http.StatusOK, true)
	return c.JSON(http.StatusCreated, res)
}

func (h *Handler) CatchPokemon(c echo.Context) error {

	result, err := h.service.CatchPokemon()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	resp := new(response.CatchPokemon)
	resp.Probability = result.Probability
	resp.Success = result.Success
	resp.Information = result.Information

	res := new(common.DefaultResponse)
	res.SetResponseData(http.StatusText(http.StatusOK), resp, http.StatusOK, true)
	return c.JSON(http.StatusAlreadyReported, res)
}

func (h *Handler) Update(c echo.Context) error {
	req := new(request.RequestUpdate)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	errVal := validation.ValidateReq(req)
	if errVal != nil {
		return c.JSON(http.StatusBadRequest, errVal)
	}

	pokemon := new(domain.Pokemon)
	pokemon.ID = c.Param("id")
	pokemon.PokemonName = req.PokemonName

	result, err := h.service.Update(pokemon)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	resp := new(response.UpdatePokemon)
	resp.ID = result.ID
	resp.PokemonName = result.PokemonName
	resp.CreatedAt = result.CreatedAt
	resp.UpdatedAt = result.UpdatedAt

	res := new(common.DefaultResponse)
	res.SetResponseData(http.StatusText(http.StatusOK), resp, http.StatusOK, true)
	return c.JSON(http.StatusOK, res)
}

func (h *Handler) List(c echo.Context) error {
	req := new(request.RequestList)

	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	results, err := h.service.List(req.PokemonName)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	res := response.NewResponseList(http.StatusText(http.StatusOK), results, http.StatusOK)
	return c.JSON(http.StatusOK, res)
}
