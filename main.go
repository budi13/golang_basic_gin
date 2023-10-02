package main

import (
	"golang_basic_gin/config"
	"golang_basic_gin/middlewares"
	"golang_basic_gin/routes"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	// fmt.Println("Framework GIN")

	config.InitDB()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/home", getHome)

	// api/v1/departments

	api := r.Group("/api/v1")
	{
		user := api.Group("/user")
		{
			user.POST("/register", routes.RegisterUser)
			user.POST("/login", routes.GenerateToken)
		}

		department := api.Group("/departments").Use(middlewares.Auth())
		{
			department.GET("/", routes.GetDepartment)
			department.GET("/:id", routes.GetDepartmentById)
			department.POST("/", routes.PostDepartment)
			department.PUT("/:id", routes.PutDepartment)
			department.DELETE("/:id", routes.DeleteDepartment)
		}

		position := api.Group("/positions").Use(middlewares.Auth())
		{
			position.GET("/", routes.GetPosition)
			position.GET("/:id", routes.GetPositionById)
			position.POST("/", routes.PostPosition)
			position.PUT("/:id", routes.PutPosition)
			position.DELETE("/:id", routes.DeletePosition)
		}

		employee := api.Group("/employees").Use(middlewares.Auth())
		{
			employee.GET("/", routes.GetEmployees)
			employee.GET("/:id", routes.GetEmployeesById)
			employee.POST("/", routes.PostEmployees)
			employee.PUT("/:id", routes.PutEmployees)
			employee.DELETE("/:id", routes.DeleteEmployees)
		}

		Inventory := api.Group("/inventories").Use(middlewares.Auth())
		{
			Inventory.GET("/", routes.GetInventories)
			Inventory.GET("/:id", routes.GetInventoriesById)
			Inventory.POST("/", routes.PostInventories)
			Inventory.PUT("/:id", routes.PutInventories)
			Inventory.DELETE("/:id", routes.DeleteInventories)
		}

		Rental := api.Group("/rentals").Use(middlewares.Auth())
		{
			Rental.GET("/", routes.GetRental)
			Rental.GET("/inventory/:id", routes.GetRentalByInventoryID)
			Rental.POST("/employee", routes.RentalByEmployeeId)
			// Inventory.PUT("/:id", routes.PutInventories)
			// Inventory.DELETE("/:id", routes.DeleteInventories)
		}

	}

	// r.GET("/department", routes.GetDepartment)
	// r.GET("/department/:id", routes.GetDepartmentById)
	// r.POST("/department", routes.PostDepartment)
	// r.PUT("/department/:id", routes.PutDepartment)
	// r.DELETE("/department/:id", routes.DeleteDepartment)

	// //Best practice pembuatan URL API pakai kalimat jamak ('s')
	// r.GET("/positions", routes.GetPosition)
	// r.GET("/positions/:id", routes.GetPositionById)
	// r.POST("/positions", routes.PostPosition)
	// r.PUT("/positions/:id", routes.PutPosition)
	// r.DELETE("/positions/:id", routes.DeletePosition)

	// r.GET("/employees", routes.GetEmployees)
	// r.GET("/employees/:id", routes.GetEmployeesById)
	// r.POST("/employees", routes.PostEmployees)
	// r.PUT("/employees/:id", routes.PutEmployees)
	// r.DELETE("/employees/:id", routes.DeleteEmployees)

	r.Run() // Listen and serve on 0.0.0.0:8080
}

func getHome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome Home",
	})
}
