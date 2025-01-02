package main

import (
	"latihanotp/config"
	"latihanotp/controllers"
	"latihanotp/repositories"
	"latihanotp/services"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.CreateDatabase()
	reposAuntentikasi := repositories.NewRepoAutentikasi(db)
	ServicAuntentikasi := services.NewServicesAuntentikasi(reposAuntentikasi)
	ControlAuntentikasi := controllers.NewControlAuntentikasi(ServicAuntentikasi)
	r := gin.Default()
	// r.Use(middlewares.JWTAuth())
	r.POST("/register", ControlAuntentikasi.Registers)
	r.POST("/otp", ControlAuntentikasi.VerifysOtp)
	r.POST("/login", ControlAuntentikasi.LoginUser)
	r.Run(":3000")
}
