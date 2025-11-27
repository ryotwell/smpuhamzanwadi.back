package api

import (
	"net/http"
	"project_sdu/model"
	"project_sdu/service"

	"github.com/gin-gonic/gin"
)

type DashboardAPI interface {
	GetDashboard(c *gin.Context)
}

type dashboardAPI struct {
	dashboardService service.DashboardService
}

func NewDashboardAPI(dashboardService service.DashboardService) *dashboardAPI {
	return &dashboardAPI{dashboardService}
}

// ====================
// GET DASHBOARD DATA
// ====================
func (d *dashboardAPI) GetDashboard(c *gin.Context) {
	totalStudents, err := d.dashboardService.GetTotalStudents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Success: false,
			Status:  http.StatusInternalServerError,
			Message: "Failed to count students",
			Errors:  map[string]string{"server": err.Error()},
		})
		return
	}

	totalPosts, err := d.dashboardService.GetTotalPosts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Success: false,
			Status:  http.StatusInternalServerError,
			Message: "Failed to count posts",
			Errors:  map[string]string{"server": err.Error()},
		})
		return
	}

	totalStudentsActiveBatch, _ := d.dashboardService.GetTotalStudentsActiveBatch()
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, model.ErrorResponse{
	// 		Success: false,
	// 		Status:  http.StatusInternalServerError,
	// 		Message: "Failed to count students from active batch",
	// 		Errors:  map[string]string{"server": err.Error()},
	// 	})
	// 	return
	// }

	totalBatch, err := d.dashboardService.GetTotalBatch()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Success: false,
			Status:  http.StatusInternalServerError,
			Message: "Failed to count batch",
			Errors:  map[string]string{"server": err.Error()},
		})
		return
	}

	activeBatch, _ := d.dashboardService.GetActiveBatch()
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, model.ErrorResponse{
	// 		Success: false,
	// 		Status:  http.StatusInternalServerError,
	// 		Message: "Failed to get active batch",
	// 		Errors:  map[string]string{"server": err.Error()},
	// 	})
	// 	return
	// }

	// Response Data
	data := gin.H{
		"total_students":              totalStudents,
		"total_posts":                 totalPosts,
		"total_students_active_batch": totalStudentsActiveBatch,
		"total_batch":                 totalBatch,
		"active_batch":                activeBatch,
	}

	c.JSON(http.StatusOK, model.SuccessResponse{
		Success: true,
		Status:  http.StatusOK,
		Message: "Dashboard data retrieved successfully",
		Data:    data,
	})
}
