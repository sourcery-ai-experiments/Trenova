package handlers

import (
	"github.com/emoss08/trenova/internal/api"
	"github.com/emoss08/trenova/internal/api/services"
	"github.com/emoss08/trenova/internal/util"
	"github.com/emoss08/trenova/internal/util/types"
	"github.com/gofiber/fiber/v2"
)

type USStateHandler struct {
	Server  *api.Server
	Service *services.USStateService
}

func NewUSStateHandler(s *api.Server) *USStateHandler {
	return &USStateHandler{
		Server:  s,
		Service: services.NewUSStateService(s),
	}
}

// GetUSStates is a handler that returns a list of states in the US.
//
// GET /us-states
func (h *USStateHandler) GetUSStates() fiber.Handler {
	return func(c *fiber.Ctx) error {
		entities, err := h.Service.GetUSStates(c.UserContext())
		if err != nil {
			errorResponse := util.CreateDBErrorResponse(err)
			return c.Status(fiber.StatusInternalServerError).JSON(errorResponse)
		}

		return c.Status(fiber.StatusOK).JSON(types.HTTPResponse{
			Results: entities,
		})
	}
}
