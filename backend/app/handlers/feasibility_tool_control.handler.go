package handlers

import (
	"net/http"

	"trenova/app/middleware"
	"trenova/app/models"
	"trenova/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetFeasibilityToolControl(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		orgID, ok := r.Context().Value(middleware.ContextKeyOrgID).(uuid.UUID)

		if !ok {
			utils.ResponseWithError(w, http.StatusInternalServerError, models.ValidationErrorDetail{
				Code:   "internalError",
				Detail: "Organization ID not found in the request context",
				Attr:   "organizationId",
			})
		}

		var feasibilityToolControl models.FeasibilityToolControl
		if err := db.
			Model(&models.FeasibilityToolControl{}).
			Where("organization_id = ?", orgID).First(&feasibilityToolControl).Error; err != nil {
			errorResponse := utils.CreateDBErrorResponse(err)
			utils.ResponseWithError(w, http.StatusInternalServerError, errorResponse)

			return
		}

		utils.ResponseWithJSON(w, http.StatusOK, feasibilityToolControl)
	}
}

func UpdateFeasibilityToolControl(db *gorm.DB, validator *utils.Validator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		orgID, ok := r.Context().Value(middleware.ContextKeyOrgID).(uuid.UUID)

		if !ok {
			utils.ResponseWithError(w, http.StatusInternalServerError, models.ValidationErrorDetail{
				Code:   "internalError",
				Detail: "Organization ID not found in the request context",
				Attr:   "organizationId",
			})
		}

		buID, ok := r.Context().Value(middleware.ContextKeyBuID).(uuid.UUID)

		if !ok {
			utils.ResponseWithError(w, http.StatusInternalServerError, models.ValidationErrorDetail{
				Code:   "internalError",
				Detail: "Business Unit ID not found in the request context",
				Attr:   "businessUnitId",
			})
		}

		var feasibilityToolControl models.FeasibilityToolControl
		if err := db.
			Where("organization_id = ? AND business_unit_id = ?", orgID, buID).
			First(&feasibilityToolControl).Error; err != nil {
			errorResponse := utils.CreateDBErrorResponse(err)
			utils.ResponseWithError(w, http.StatusInternalServerError, errorResponse)

			return
		}

		if err := utils.ParseBodyAndValidate(validator, w, r, &feasibilityToolControl); err != nil {
			return
		}

		if err := db.Save(&feasibilityToolControl).Error; err != nil {
			errorResponse := utils.CreateDBErrorResponse(err)
			utils.ResponseWithError(w, http.StatusInternalServerError, errorResponse)

			return
		}

		utils.ResponseWithJSON(w, http.StatusOK, feasibilityToolControl)
	}
}
