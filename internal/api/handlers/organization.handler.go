package handlers

import (
	"github.com/emoss08/trenova/internal/api"
	"github.com/emoss08/trenova/internal/api/services"
	"github.com/emoss08/trenova/internal/util"
	"github.com/emoss08/trenova/internal/util/types"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// GetUserOrganization is a handler that returns the organization of the currently authenticated user.
func GetUserOrganization(s *api.Server) fiber.Handler {
	return func(c *fiber.Ctx) error {
		buID, buOK := c.Locals(util.CTXBusinessUnitID).(uuid.UUID)
		orgID, orgOK := c.Locals(util.CTXOrganizationID).(uuid.UUID)

		if !buOK || !orgOK {
			return c.Status(fiber.StatusInternalServerError).JSON(types.ValidationErrorResponse{
				Type: "internalServerError",
				Errors: []types.ValidationErrorDetail{
					{
						Code:   "internalServerError",
						Detail: "Internal server error",
						Attr:   "session",
					},
				},
			},
			)
		}

		user, err := services.NewOrganizationOps(s).GetUserOrganization(c.UserContext(), buID, orgID)
		if err != nil {
			errorResponse := util.CreateDBErrorResponse(err)
			return c.Status(fiber.StatusInternalServerError).JSON(errorResponse)
		}

		return c.Status(fiber.StatusOK).JSON(user)
	}
}
