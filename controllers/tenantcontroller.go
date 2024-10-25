package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/retere/IOTSmartMeasureKWH/entity/tenantentity"
	"github.com/retere/IOTSmartMeasureKWH/helpers"
	"github.com/retere/IOTSmartMeasureKWH/repository"
	"net/http"
)

func CreateNewTenant(c *gin.Context) {
	var input repository.Tenants

	// Validate the incoming JSON payload
	if err := c.ShouldBindJSON(&input); err != nil {
		response := &helpers.APIFailure{
			API: &helpers.API{
				Status:  "failed",
				State:   "non active",
				Message: "Invalid request payload",
			},
		}

		helpers.FailureResponseJSON(c.Writer, http.StatusBadRequest, response)
		return
	}

	savedTenant, err := input.Save()
	if err != nil {
		helpers.ERROR(c.Writer, http.StatusInternalServerError, err)
		return
	}

	response := &helpers.APISuccess{
		API: &helpers.API{
			Status:  "failed",
			State:   "non active",
			Message: "Tenant saved successfully",
		},
		Meta: nil,
		Data: savedTenant,
	}

	helpers.SuccessResponseJSON(c.Writer, http.StatusCreated, response)
	return
}

func UpdateTenantByID(c *gin.Context) {
	var input tenantentity.UpdateTenant
	tenantID := c.Param("id") // Get tenant ID from URL parameter

	// Bind the incoming JSON payload to the TenantRequest struct
	if err := c.ShouldBindJSON(&input); err != nil {
		response := &helpers.APIFailure{
			API: &helpers.API{
				Status:  "failed",
				State:   "non active",
				Message: "Invalid request payload",
			},
		}

		helpers.FailureResponseJSON(c.Writer, http.StatusBadRequest, response)
		return
	}

	// Find the tenant by TenantID
	var existingTenant repository.Tenants
	if err := repository.FindTenantByID(tenantID, &existingTenant); err != nil {
		response := &helpers.APIFailure{
			API: &helpers.API{
				Status:  "failed",
				State:   "non active",
				Message: "Tenant Not Found",
			},
		}

		helpers.FailureResponseJSON(c.Writer, http.StatusBadRequest, response)
		return
	}

	// Update tenant fields (only the fields that are non-zero in the request)
	existingTenant.TenantName = input.TenantName
	existingTenant.Address = input.Address
	existingTenant.UpdatedAt = input.UpdatedAt // Set UpdatedAt to the new time (or use time.Now())

	// Save the updated tenant
	if err := repository.SaveExistingTenant(&existingTenant, tenantID); err != nil {
		helpers.ERROR(c.Writer, http.StatusInternalServerError, err)
		return
	}

	response := &helpers.APISuccess{
		API: &helpers.API{
			Status:  "success",
			State:   "active",
			Message: "Tenant saved successfully",
		},
		Meta: nil,
		Data: existingTenant,
	}

	helpers.SuccessResponseJSON(c.Writer, http.StatusCreated, response)
	return
}
