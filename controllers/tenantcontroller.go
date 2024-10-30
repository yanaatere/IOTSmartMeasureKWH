package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/retere/IOTSmartMeasureKWH/entity/tenantentity"
	"github.com/retere/IOTSmartMeasureKWH/helpers"
	"github.com/retere/IOTSmartMeasureKWH/repository"
	"net/http"
	"strconv"
)

func CreateNewTenant(c *gin.Context) {
	var input tenantentity.UpdateTenant

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

	err := repository.SaveTenants(&input)
	if err != nil {
		helpers.ERROR(c.Writer, http.StatusInternalServerError, err)
		return
	}

	response := &helpers.APISuccess{
		API: &helpers.API{
			Status:  "Success",
			State:   "Active",
			Message: "Tenant saved successfully",
		},
		Meta: nil,
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
	var existingTenant tenantentity.Tenants
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

	existingTenant.TenantName = input.TenantName
	existingTenant.Address = input.Address

	// SaveTenants the updated tenant
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

func GetTenantByID(c *gin.Context) {
	tenantID := c.Param("id")

	var tenants tenantentity.Tenants
	err := repository.FindTenantByID(tenantID, &tenants)
	if err != nil {
		response := &helpers.APIFailure{
			API: &helpers.API{
				Status:  "failed",
				State:   "non active",
				Message: "Tenant Not Found",
			},
		}
		helpers.FailureResponseJSON(c.Writer, http.StatusNotFound, response)
	}

	response := &helpers.APISuccess{
		API: &helpers.API{
			Status:  "success",
			State:   "active",
			Message: "FInd One Tenant",
		},
		Meta: nil,
		Data: tenants,
	}

	helpers.SuccessResponseJSON(c.Writer, http.StatusCreated, response)
}

func DeleteTenantByID(c *gin.Context) {
	tenantID := c.Param("id")

	var tenants repository.Tenants
	err := repository.DeleteById(tenantID, &tenants)
	if err != nil {
		response := &helpers.APIFailure{
			API: &helpers.API{
				Status:  "failed",
				State:   "non active",
				Message: "Failed To Delete TenantId",
			},
		}
		helpers.FailureResponseJSON(c.Writer, http.StatusNotFound, response)
	}

	response := &helpers.APISuccess{
		API: &helpers.API{
			Status:  "success",
			State:   "active",
			Message: "Successfully deleted",
		},
		Meta: nil,
		Data: tenants,
	}

	helpers.SuccessResponseJSON(c.Writer, http.StatusCreated, response)
}

func GetAllTenants(c *gin.Context) {
	pageParam := c.Param("page")
	sizeParam := c.Param("size")

	page, err := strconv.Atoi(pageParam)
	if err != nil || page < 1 {
		page = 1
	}

	size, err := strconv.Atoi(sizeParam)
	if err != nil || size < 1 {
		size = 10
	}

	tenants, err := repository.FindAllTenants(page, size)
	if err != nil {
		response := &helpers.APIFailure{
			API: &helpers.API{
				Status:  "failed",
				State:   "non active",
				Message: "Failed To Delete TenantId",
			},
		}
		helpers.FailureResponseJSON(c.Writer, http.StatusNotFound, response)
	}

	response := &helpers.APISuccess{
		API: &helpers.API{
			Status:  "success",
			State:   "active",
			Message: "Success Get All Tenatants",
		},
		Meta: nil,
		Data: tenants,
	}

	helpers.SuccessResponseJSON(c.Writer, http.StatusCreated, response)

}
