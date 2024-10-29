package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/retere/IOTSmartMeasureKWH/controllers"
)

func tenantGroupRouter(baseRouter *gin.RouterGroup) {
	tenant := baseRouter.Group("/tenant")

	tenant.POST("/all/page/:page/size/:size", controllers.GetAllTenants)
	tenant.GET("/get/:id", controllers.GetTenantByID)
	tenant.POST("/create", controllers.CreateNewTenant)
	tenant.PATCH("/update/:id", controllers.UpdateTenantByID)
	tenant.DELETE("/delete/:id", controllers.DeleteTenantByID)
}

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	versionRouter := r.Group("/api/v1")
	tenantGroupRouter(versionRouter)

	return r
}
