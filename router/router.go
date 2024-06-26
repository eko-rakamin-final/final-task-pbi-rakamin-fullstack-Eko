package router

import (
	"rakamin-go/controllers"
	"rakamin-go/database"
	"rakamin-go/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()

    db := database.ConnectDB() 

    userController := controllers.UserController{DB: db}
    r.POST("/users/register", userController.RegisterUser)
    r.POST("/users/login", userController.LoginUser)
    r.PUT("/users/:id", userController.UpdateUser)
    r.DELETE("/users/:id", userController.DeleteUser)
	
	photoController := controllers.PhotoController{DB: db}

    photoRoutes := r.Group("/photos")
    photoRoutes.Use(middlewares.AuthMiddleware())
    {
        photoRoutes.GET("/", photoController.GetPhotos)
        photoRoutes.POST("/", photoController.UploadPhoto)
        photoRoutes.PUT("/:id", photoController.UpdatePhoto)
        photoRoutes.DELETE("/:id", photoController.DeletePhoto)
    }


    return r
}