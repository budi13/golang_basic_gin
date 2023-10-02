package routes

import (
	"golang_basic_gin/config"
	"golang_basic_gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func GetEmployees(c *gin.Context) {
	employee := []models.Employee{}

	// dengan relational db
	config.DB.Preload(clause.Associations).Find(&employee)

	getEmployeeResponse := []models.GetEmployeeResponse{}

	for _, e := range employee {

		em := models.GetEmployeeResponse{
			ID:         e.ID,
			Name:       e.Name,
			Address:    e.Address,
			Email:      e.Email,
			PositionID: e.PositionID,
			Position: models.PositionResponse{
				ID:   e.Position.ID,
				Name: e.Position.Name,
				Code: e.Position.Code,
			},
		}

		getEmployeeResponse = append(getEmployeeResponse, em)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully retrieved department",
		"data":    getEmployeeResponse,
	})
}

func GetEmployeesById(c *gin.Context) {
	id := c.Param("id")

	var employee models.Employee

	// dengan relational db
	data := config.DB.Preload(clause.Associations).First(&employee, "id = ?", id)

	if data.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "404 Data Not Found",
			"message": "Employee not found",
		})
		return
	}

	em := models.GetEmployeeResponse{
		ID:         employee.ID,
		Name:       employee.Name,
		Address:    employee.Address,
		Email:      employee.Email,
		PositionID: employee.PositionID,
		Position: models.PositionResponse{
			ID:   employee.Position.ID,
			Name: employee.Position.Name,
			Code: employee.Position.Code,
		},
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully retrieved employee",
		"data":    em,
	})
}

func PostEmployees(c *gin.Context) {

	// ambil data post dari JSON
	var employees models.Employee
	c.BindJSON(&employees)

	// insert data to DB
	config.DB.Create(&employees)

	c.JSON(http.StatusCreated, gin.H{
		"message": "Successfully created employee",
		"data":    employees,
	})
}

func PutEmployees(c *gin.Context) {
	id := c.Param("id")

	var employees models.Employee

	var reqEmployee models.Employee
	c.BindJSON(&reqEmployee)

	config.DB.Model(&employees).Where("id = ?", id).Updates(reqEmployee)

	c.JSON(200, gin.H{
		"message": "Update Success",
		"data":    employees,
	})
}

func DeleteEmployees(c *gin.Context) {
	id := c.Param("id")

	var employees models.Employee
	// data := config.DB.Where("id = ?", id).Find(&department)
	data := config.DB.First(&employees, "id = ?", id)

	if data.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "Data not found",
			"message": "Data not found",
		})
		return
	}

	config.DB.Delete(&employees, id)

	c.JSON(200, gin.H{
		"message": "Delete Success",
	})
}
