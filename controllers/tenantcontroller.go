package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/retere/IOTSmartMeasureKWH/entity/tenantentity"
	"github.com/retere/IOTSmartMeasureKWH/repository"
	"net/http"
)

func CreateNewTenant(c *gin.Context) {
	var input repository.Tenants

	// Validate the incoming JSON payload
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": "Invalid request payload",
			"error":   err.Error(),
			"data":    nil,
		})
		return
	}

	savedTenant, err := input.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "failed",
			"message": "Failed to save tenant",
			"error":   err.Error(),
			"data":    nil,
		})
		return
	}

	// Respond with success message if save operation succeeds
	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "Tenant saved successfully",
		"data":    savedTenant,
	})
}

func UpdateTenantByID(c *gin.Context) {
	var input tenantentity.UpdateTenant
	tenantID := c.Param("id") // Get tenant ID from URL parameter

	// Bind the incoming JSON payload to the TenantRequest struct
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": "Invalid request payload",
			"error":   err.Error(),
			"data":    nil,
		})
		return
	}

	// Find the tenant by TenantID
	var existingTenant repository.Tenants
	if err := repository.FindTenantByID(tenantID, &existingTenant); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "failed",
			"message": "Tenant not found",
			"error":   err.Error(),
			"data":    nil,
		})
		return
	}

	// Update tenant fields (only the fields that are non-zero in the request)
	existingTenant.TenantName = input.TenantName
	existingTenant.Address = input.Address
	existingTenant.UpdatedAt = input.UpdatedAt // Set UpdatedAt to the new time (or use time.Now())

	// Save the updated tenant
	if err := repository.SaveExistingTenant(&existingTenant, tenantID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "failed",
			"message": "Failed to update tenant",
			"error":   err.Error(),
			"data":    nil,
		})
		return
	}

	// Respond with success message and updated tenant
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Tenant updated successfully",
		"data":    existingTenant,
	})
}
