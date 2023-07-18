package healthcheck

import (
	"net/http"

	port "pokemaster-api/core/port/healthCheck"
	"pokemaster-api/interface/api/common"

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

// Health godoc
// @Summary 	Health check
// @Description Health check
// @Tags 		health-check
// @Accept 		json
// @Produce 	json
// @Success 	200 	{object}	common.DefaultResponse
// @Failure     500		{object}	common.DefaultResponse
// @Router /api/v1/health [get]
func (h *Handler) Get(c echo.Context) error {
	get := h.service.Get()
	response := new(common.DefaultResponse)
	response.SetResponseData(http.StatusText(http.StatusOK), get, http.StatusOK, true)
	return c.JSON(http.StatusOK, response)
}
