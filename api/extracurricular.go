package api

import (
	"net/http"
	"project_sdu/model"
	"project_sdu/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ExtracurricularAPI interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	GetByID(c *gin.Context)
	GetAll(c *gin.Context)
}

type extracurricularAPI struct {
	exService service.ExtracurricularService
}

func NewExtracurricularAPI(exService service.ExtracurricularService) *extracurricularAPI {
	return &extracurricularAPI{exService}
}

// ====================
// CREATE
// ====================
func (e *extracurricularAPI) Create(c *gin.Context) {
	var ex model.Extracurricular
	if err := c.ShouldBindJSON(&ex); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Success: false,
			Status:  http.StatusBadRequest,
			Message: "Validation failed",
			Errors:  map[string]string{"body": "Invalid JSON format"},
		})
		return
	}

	if ex.Name == "" {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Success: false,
			Status:  http.StatusBadRequest,
			Message: "Validation failed",
			Errors:  map[string]string{"name": "name is required"},
		})
		return
	}

	if err := e.exService.Create(&ex); err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Success: false,
			Status:  http.StatusInternalServerError,
			Message: "Failed to create extracurricular",
			Errors:  map[string]string{"server": err.Error()},
		})
		return
	}

	c.JSON(http.StatusCreated, model.SuccessResponse{
		Success: true,
		Status:  http.StatusCreated,
		Message: "Extracurricular created successfully",
		Data:    ex,
	})
}

// ====================
// UPDATE
// ====================
func (e *extracurricularAPI) Update(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Success: false,
			Status:  http.StatusBadRequest,
			Message: "Invalid ID",
		})
		return
	}

	var ex model.Extracurricular
	if err := c.ShouldBindJSON(&ex); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Success: false,
			Status:  http.StatusBadRequest,
			Message: "Invalid JSON format",
			Errors:  map[string]string{"body": err.Error()},
		})
		return
	}

	if err := e.exService.Update(id, &ex); err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Success: false,
			Status:  http.StatusInternalServerError,
			Message: "Failed to update extracurricular",
			Errors:  map[string]string{"server": err.Error()},
		})
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse{
		Success: true,
		Status:  http.StatusOK,
		Message: "Extracurricular updated successfully",
		Data:    ex,
	})
}

// ====================
// DELETE
// ====================
func (e *extracurricularAPI) Delete(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Success: false,
			Status:  http.StatusBadRequest,
			Message: "Invalid ID",
		})
		return
	}

	if err := e.exService.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Success: false,
			Status:  http.StatusInternalServerError,
			Message: "Failed to delete extracurricular",
			Errors:  map[string]string{"server": err.Error()},
		})
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse{
		Success: true,
		Status:  http.StatusOK,
		Message: "Extracurricular deleted successfully",
	})
}

// ====================
// GET BY ID
// ====================
func (e *extracurricularAPI) GetByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Success: false,
			Status:  http.StatusBadRequest,
			Message: "Invalid ID",
		})
		return
	}

	ex, err := e.exService.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, model.ErrorResponse{
			Success: false,
			Status:  http.StatusNotFound,
			Message: "Extracurricular not found",
			Errors:  map[string]string{"server": err.Error()},
		})
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse{
		Success: true,
		Status:  http.StatusOK,
		Message: "Extracurricular retrieved successfully",
		Data:    ex,
	})
}

// ====================
// GET ALL
// ====================
func (e *extracurricularAPI) GetAll(c *gin.Context) {
	limitParam := c.DefaultQuery("limit", "10")
	pageParam := c.DefaultQuery("page", "1")

	limit, _ := strconv.Atoi(limitParam)
	page, _ := strconv.Atoi(pageParam)

	data, err := e.exService.GetAll(limit, page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Success: false,
			Status:  http.StatusInternalServerError,
			Message: "Failed to retrieve extracurricular",
			Errors:  map[string]string{"server": err.Error()},
		})
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse{
		Success: true,
		Status:  http.StatusOK,
		Message: "Extracurricular retrieved successfully",
		Data:    data,
		Meta: gin.H{
			"page":  page,
			"limit": limit,
		},
	})
}
