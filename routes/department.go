package routes

import (
	"golang_basic_gin/config"
	"golang_basic_gin/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// func GetDepartment(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{
// 		"meesage": "Successfully retrieved department",
// 		"data":    "data",
// 	})
// }

func GetDepartment(c *gin.Context) {
	departments := []models.Department{}

	// tanpa relational db
	// config.DB.Find(&departments)

	// dengan relational db
	config.DB.Preload("Positions").Find(&departments)

	GetDepartmentResponse := []models.GetDepartmentResponse{}

	for _, d := range departments {

		positions := []models.PositionResponse{}
		for _, p := range d.Positions {
			pos := models.PositionResponse{
				ID:   p.ID,
				Name: p.Name,
				Code: p.Code,
			}

			positions = append(positions, pos)
		}

		dept := models.GetDepartmentResponse{
			ID:        d.ID,
			Name:      d.Name,
			Code:      d.Code,
			Positions: positions,
		}

		GetDepartmentResponse = append(GetDepartmentResponse, dept)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully retrieved department",
		"data":    GetDepartmentResponse,
	})
}

func GetDepartmentById(c *gin.Context) {
	id := c.Param("id")

	var department models.Department
	// data := config.DB.Where("id = ?", id).Find(&department)

	// tanpa relational db
	// data := config.DB.First(&department, "id = ?", id)

	// dengan relational db
	data := config.DB.Preload("Positions").First(&department, "id = ?", id)

	if data.Error != nil {
		log.Printf(data.Error.Error())
		c.JSON(http.StatusNotFound, gin.H{
			"message": "department not found",
		})
		return
	}

	positions := []models.PositionResponse{}
	for _, p := range department.Positions {
		pos := models.PositionResponse{
			ID:   p.ID,
			Name: p.Name,
			Code: p.Code,
		}
		positions = append(positions, pos)
	}

	getDepartmentResponse := models.GetDepartmentResponse{
		ID:        department.ID,
		Name:      department.Name,
		Code:      department.Code,
		Positions: positions,
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully retrieved department",
		"data":    getDepartmentResponse,
	})
}

func PostDepartment(c *gin.Context) {

	// ambil data post dari JSON
	var department models.Department
	c.BindJSON(&department)

	// insert data to DB
	config.DB.Create(&department)

	c.JSON(http.StatusCreated, gin.H{
		"message": "Successfully created department",
		"data":    department,
	})
}

func PutDepartment(c *gin.Context) {
	id := c.Param("id")

	var department models.Department
	// data := config.DB.Where("id = ?", id).Find(&department)
	data := config.DB.First(&department, "id = ?", id)

	if data.Error != nil {
		log.Printf(data.Error.Error())
		c.JSON(http.StatusNotFound, gin.H{
			"message": "department not found",
		})
		return
	}
	c.BindJSON(&department)

	config.DB.Model(&department).Where("id = ?", id).Updates(&department)

	c.JSON(http.StatusOK, gin.H{
		"message": "Update Successfully",
		"data":    department,
	})
}

func DeleteDepartment(c *gin.Context) {
	id := c.Param("id")

	var department models.Department
	// data := config.DB.Where("id = ?", id).Find(&department)
	data := config.DB.First(&department, "id = ?", id)

	if data.Error != nil {
		log.Printf(data.Error.Error())
		c.JSON(http.StatusNotFound, gin.H{
			"message": "department not found",
		})
		return
	}

	config.DB.Delete(&department, id)

	c.JSON(http.StatusOK, gin.H{
		"message": "Delete Successfully",
	})
}
