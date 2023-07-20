package handler

import (
	"net/http"
	domain "pokemaster-api/core/domain/user"
	port "pokemaster-api/core/port/user"
	"pokemaster-api/interface/api/common"
	"pokemaster-api/interface/api/extl/v1/user/request"
	"pokemaster-api/interface/api/extl/v1/user/response"
	"pokemaster-api/interface/api/utils/validation"

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

	pokemon := new(domain.User)
	pokemon.FullName = req.FullName
	pokemon.Email = req.Email
	pokemon.Password = req.Password

	err := h.service.Insert(pokemon)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	res := new(common.DefaultResponseNoData)
	res.SetResponseDataNoData(http.StatusText(http.StatusCreated), http.StatusCreated, true)
	return c.JSON(http.StatusCreated, res)
}

func (h *Handler) Get(c echo.Context) error {

	results, err := h.service.List()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	res := response.NewResponseList(http.StatusText(http.StatusOK), results, http.StatusOK)
	return c.JSON(http.StatusOK, res)
}
