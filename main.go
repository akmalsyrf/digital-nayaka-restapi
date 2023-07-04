package main

import (
	"fmt"
	"log"
	"os"

	annualreviews "digital-nayaka-test/annualReviews"
	"digital-nayaka-test/employee"
	"digital-nayaka-test/handler"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln(err)
	}

	DB_HOST := os.Getenv("DB_HOST")
	DB_USER := os.Getenv("DB_USER")
	DB_PASS := os.Getenv("DB_PASS")
	DB_NAME := os.Getenv("DB_NAME")
	DB_PORT := os.Getenv("DB_PORT")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", DB_HOST, DB_USER, DB_PASS, DB_NAME, DB_PORT)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Connected to postgres")

	if err := db.AutoMigrate(&employee.Employee{}, &annualreviews.AnnualReviews{}); err != nil {
		log.Fatalln(err)
	}

	router := gin.Default()
	router.Use(cors.Default())
	api := router.Group("/api/v1")

	employeeRepository := employee.NewRepository(db)
	employeeService := employee.NewService(employeeRepository)
	employeeHandler := handler.NewEmployeeHandler(employeeService)

	api.GET("/employees", employeeHandler.GetEmployees)
	router.Run(":8080")
}
