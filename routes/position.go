package routes

import (
	"golang_basic_gin/config"
	"golang_basic_gin/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPosition(c *gin.Context) {
	positions := []models.Position{}

	// tanpa relational db
	// config.DB.Find(&positions)

	// dengan relational db
	config.DB.Preload("Department").Find(&positions)

	getPositionResponse := []models.GetPositionResponse{}

	for _, p := range positions {
		department := models.DepartmentResponse{
			ID:   p.Department.ID,
			Name: p.Department.Name,
			Code: p.Department.Code,
		}

		post := models.GetPositionResponse{
			ID:           p.ID,
			Name:         p.Name,
			Code:         p.Code,
			DepartmentID: p.DepartmentID,
			Department:   department,
		}

		getPositionResponse = append(getPositionResponse, post)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully retrieved position",
		"data":    getPositionResponse,
	})
}

func GetPositionById(c *gin.Context) {
	id := c.Param("id")

	var position models.Position
	// data := config.DB.Where("id = ?", id).Find(&department)

	// tanpa relational db
	// data := config.DB.First(&position, "id = ?", id)

	// dengan relational db
	data := config.DB.Preload("Department").First(&position, "id = ?", id)

	if data.Error != nil {
		log.Printf(data.Error.Error())
		c.JSON(http.StatusNotFound, gin.H{
			"message": "position not found",
		})
		return
	}
	dept := models.DepartmentResponse{
		ID:   position.Department.ID,
		Name: position.Department.Name,
		Code: position.Department.Code,
	}

	getPositionResponse := models.GetPositionResponse{
		ID:           position.ID,
		Name:         position.Name,
		Code:         position.Code,
		DepartmentID: position.DepartmentID,
		Department:   dept,
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully retrieved position",
		"data":    getPositionResponse,
	})
}

func PostPosition(c *gin.Context) {

	// ambil data post dari JSON
	var position models.Position
	c.BindJSON(&position)

	// insert data to DB
	config.DB.Create(&position)

	c.JSON(http.StatusCreated, gin.H{
		"message": "Successfully created position",
		"data":    position,
	})
}

func PutPosition(c *gin.Context) {
	id := c.Param("id")

	var position models.Position
	// data := config.DB.Where("id = ?", id).Find(&department)
	data := config.DB.First(&position, "id = ?", id)

	if data.Error != nil {
		log.Printf(data.Error.Error())
		c.JSON(http.StatusNotFound, gin.H{
			"message": "position not found",
		})
		return
	}
	c.BindJSON(&position)

	config.DB.Model(&position).Where("id = ?", id).Updates(&position)

	c.JSON(http.StatusOK, gin.H{
		"message": "Update Successfully",
		"data":    position,
	})
}

func DeletePosition(c *gin.Context) {
	id := c.Param("id")

	var position models.Position
	// data := config.DB.Where("id = ?", id).Find(&department)
	data := config.DB.First(&position, "id = ?", id)

	if data.Error != nil {
		log.Printf(data.Error.Error())
		c.JSON(http.StatusNotFound, gin.H{
			"message": "position not found",
		})
		return
	}

	config.DB.Delete(&position, id)

	c.JSON(http.StatusOK, gin.H{
		"message": "Delete Successfully",
	})
}
