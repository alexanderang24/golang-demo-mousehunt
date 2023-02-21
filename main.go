package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang-demo-mousehunt/controllers"
	"golang-demo-mousehunt/database"
	"golang-demo-mousehunt/middleware"
	"golang-demo-mousehunt/repository"
	"golang-demo-mousehunt/services"
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
		os.Getenv("PGDATABASE"),
	)

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

	// trap
	trapRepo := repository.NewTrapRepository(db)
	trapService := services.NewTrapService(trapRepo)
	trapController := controllers.NewTrapController(trapService)
	trap := router.Group("/trap")
	trap.GET("", trapController.GetAllTraps)
	trap.GET("/:id", trapController.GetTrap)
	trap.POST("", middleware.VerifyJWT, trapController.InsertTrap)
	trap.PUT("/:id", middleware.VerifyJWT, trapController.UpdateTrap)
	trap.DELETE("/:id", middleware.VerifyJWT, trapController.DeleteTrap)

	// location
	locationRepo := repository.NewLocationRepository(db)
	locationService := services.NewLocationService(locationRepo)
	locationController := controllers.NewLocationController(locationService)
	location := router.Group("/location")
	location.GET("", locationController.GetAllLocations)
	location.GET("/:id", locationController.GetLocation)
	location.POST("", middleware.VerifyJWT, locationController.InsertLocation)
	location.PUT("/:id", middleware.VerifyJWT, locationController.UpdateLocation)
	location.DELETE("/:id", middleware.VerifyJWT, locationController.DeleteLocation)

	// mouse
	mouseRepo := repository.NewMouseRepository(db)
	mouseService := services.NewMouseService(mouseRepo)
	mouseController := controllers.NewMouseController(mouseService)
	mouse := router.Group("/mouse")
	mouse.GET("", mouseController.GetAllMice)
	mouse.GET("/:id", mouseController.GetMouse)
	mouse.POST("", middleware.VerifyJWT, mouseController.InsertMouse)
	mouse.PUT("/:id", middleware.VerifyJWT, mouseController.UpdateMouse)
	mouse.DELETE("/:id", middleware.VerifyJWT, mouseController.DeleteMouse)

	// user - admin
	userRepo := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)
	user := router.Group("/user")
	user.GET("", userController.GetAllUsers)
	user.GET("/:id", userController.GetUser)
	user.POST("", middleware.VerifyJWT, middleware.AdminOnly, userController.InsertUser)
	user.PUT("/:id", middleware.VerifyJWT, userController.UpdateUser)
	user.DELETE("/:id", middleware.VerifyJWT, userController.DeleteUser)


	// hunt history
	huntRepo := repository.NewHuntHistoryRepository(db)
	huntService := services.NewHistoryService(huntRepo)
	huntController := controllers.NewHuntHistoryController(huntService)
	hunt := router.Group("/hunt")

	// player APIs
	user.POST("/register", userController.Register)
	user.POST("/login", userController.Login)
	//trap.POST("/trap/:id/buy", middleware.BasicAuth, trapController.BuyTrap)
	//location.POST("/:id/travel", middleware.BasicAuth, locationController.TravelToLocation)
	hunt.GET("", huntController.GetAllHuntHistories)
	hunt.POST("", middleware.VerifyJWT, huntController.DoHunt)

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