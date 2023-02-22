package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang-demo-mousehunt/controllers"
	"golang-demo-mousehunt/database"
	"golang-demo-mousehunt/middleware"
	"os"

	_ "github.com/lib/pq"
)

var (
	db *sql.DB
	err error
)

func main() {
	loadConfig()

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("PGHOST"),
		os.Getenv("PGPORT"),
		os.Getenv("PGUSER"),
		os.Getenv("PGPASSWORD"),
		os.Getenv("PGDATABASE"))

	db, err = sql.Open("postgres", psqlInfo)
	err = db.Ping()
	if err != nil {
		fmt.Println("DB Connection Failed")
		panic(err)
	} else {
		fmt.Println("DB Connection Success")
	}

	database.DbMigrate(db)
	defer db.Close()

	router := gin.Default()

	// location
	location := router.Group("/location", middleware.VerifyJWT)
	location.GET("", controllers.GetAllLocations)
	location.GET("/:id", controllers.GetLocation)
	location.POST("", middleware.AdminOnly, controllers.InsertLocation)
	location.PUT("/:id", middleware.AdminOnly, controllers.UpdateLocation)
	location.DELETE("/:id", middleware.AdminOnly, controllers.DeleteLocation)
	location.POST("/:id/travel", controllers.TravelToLocation)

	// mouse
	mouse := router.Group("/mouse", middleware.VerifyJWT)
	mouse.GET("", middleware.AdminOnly, controllers.GetAllMice)
	mouse.GET("/:id", controllers.GetMouse)
	mouse.POST("", middleware.AdminOnly, controllers.InsertMouse)
	mouse.PUT("/:id", middleware.AdminOnly, controllers.UpdateMouse)
	mouse.DELETE("/:id", middleware.AdminOnly, controllers.DeleteMouse)

	// user
	user := router.Group("/user", middleware.VerifyJWT)
	user.GET("", middleware.AdminOnly, controllers.GetAllUsers)
	user.GET("/:id", controllers.GetUser)
	user.POST("", middleware.AdminOnly, controllers.InsertUser)
	user.PUT("/:id", middleware.AdminOnly, controllers.UpdateUser)
	user.DELETE("/:id", middleware.AdminOnly, controllers.DeleteUser)
	router.POST("/user/register", controllers.Register)
	router.POST("/user/login", controllers.Login)

	// trap
	trap := router.Group("/trap", middleware.VerifyJWT)
	trap.GET("", controllers.GetAllTraps)
	trap.GET("/:id", controllers.GetTrap)
	trap.POST("", middleware.AdminOnly, controllers.InsertTrap)
	trap.PUT("/:id", middleware.AdminOnly, controllers.UpdateTrap)
	trap.DELETE("/:id", middleware.AdminOnly, controllers.DeleteTrap)
	trap.POST("/:id/buy", controllers.BuyTrap)


	// hunt history
	hunt := router.Group("/hunt", middleware.VerifyJWT)
	hunt.GET("", controllers.GetAllHuntHistories)
	hunt.POST("", controllers.DoHunt)

	err :=router.Run(":" + os.Getenv("PORT"))
	if err != nil {
		return
	}
}

func loadConfig() {
	// .env Config
	err = godotenv.Load("config/.env")
	if err != nil {
		fmt.Println("Failed to load file environment")
	} else {
		fmt.Println("Successfully load file environment")
	}
}