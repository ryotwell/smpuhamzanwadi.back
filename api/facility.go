package api

import (
	"net/http"
	"project_sdu/model"
	"project_sdu/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FacilityAPI interface {
	CreateFacility(c *gin.Context)
	GetFacilityByID(c *gin.Context)
	GetAllFacilities(c *gin.Context)
	UpdateFacility(c *gin.Context)
	DeleteFacility(c *gin.Context)
}

type facilityAPI struct {
	facilityService service.FacilityService
}

func NewFacilityAPI(facilityService service.FacilityService) *facilityAPI {
	return &facilityAPI{facilityService}
}

// ====================
// CREATE FACILITY
// ====================
func (f *facilityAPI) CreateFacility(c *gin.Context) {
	var facility model.Facility
	if err := c.ShouldBindJSON(&facility); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Success: false,
			Status:  http.StatusBadRequest,
			Message: "Invalid JSON format",
			Errors:  map[string]string{"body": err.Error()},
		})
		return
	}

	if facility.Name == "" {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Success: false,
			Status:  http.StatusBadRequest,
			Message: "Validation failed",
			Errors:  map[string]string{"name": "Name is required"},
		})
		return
	}

	if err := f.facilityService.Create(&facility); err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Success: false,
			Status:  http.StatusInternalServerError,
			Message: "Failed to create facility",
			Errors:  map[string]string{"server": err.Error()},
		})
		return
	}

	c.JSON(http.StatusCreated, model.SuccessResponse{
		Success: true,
		Status:  http.StatusCreated,
		Message: "Facility created successfully",
		Data:    facility,
	})
}

// ====================
// GET FACILITY BY ID
// ====================
func (f *facilityAPI) GetFacilityByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Success: false,
			Status:  http.StatusBadRequest,
			Message: "Invalid facility ID",
		})
		return
	}

	facility, err := f.facilityService.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, model.ErrorResponse{
			Success: false,
			Status:  http.StatusNotFound,
			Message: "Facility not found",
			Errors:  map[string]string{"server": err.Error()},
		})
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse{
		Success: true,
		Status:  http.StatusOK,
		Message: "Facility retrieved successfully",
		Data:    facility,
	})
}

// ====================
// GET ALL FACILITIES
// ====================
func (f *facilityAPI) GetAllFacilities(c *gin.Context) {
	limitParam := c.DefaultQuery("limit", "10")
	pageParam := c.DefaultQuery("page", "1")

	limit, _ := strconv.Atoi(limitParam)
	page, _ := strconv.Atoi(pageParam)

	facilities, err := f.facilityService.GetAll(limit, page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Success: false,
			Status:  http.StatusInternalServerError,
			Message: "Failed to retrieve facilities",
			Errors:  map[string]string{"server": err.Error()},
		})
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse{
		Success: true,
		Status:  http.StatusOK,
		Message: "Facilities retrieved successfully",
		Data:    facilities,
	})
}

// ====================
// UPDATE FACILITY
// ====================
func (f *facilityAPI) UpdateFacility(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Success: false,
			Status:  http.StatusBadRequest,
			Message: "Invalid facility ID",
		})
		return
	}

	var facility model.Facility
	if err := c.ShouldBindJSON(&facility); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Success: false,
			Status:  http.StatusBadRequest,
			Message: "Invalid JSON format",
			Errors:  map[string]string{"body": err.Error()},
		})
		return
	}

	if err := f.facilityService.Update(id, &facility); err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Success: false,
			Status:  http.StatusInternalServerError,
			Message: "Failed to update facility",
			Errors:  map[string]string{"server": err.Error()},
		})
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse{
		Success: true,
		Status:  http.StatusOK,
		Message: "Facility updated successfully",
		Data:    facility,
	})
}

// ====================
// DELETE FACILITY
// ====================
func (f *facilityAPI) DeleteFacility(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Success: false,
			Status:  http.StatusBadRequest,
			Message: "Invalid facility ID",
		})
		return
	}

	if err := f.facilityService.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Success: false,
			Status:  http.StatusInternalServerError,
			Message: "Failed to delete facility",
			Errors:  map[string]string{"server": err.Error()},
		})
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse{
		Success: true,
		Status:  http.StatusOK,
		Message: "Facility deleted successfully",
	})
}
