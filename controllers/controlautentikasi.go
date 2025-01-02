package controllers

import (
	"latihanotp/middlewares"
	"latihanotp/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ControlAuntentikasi struct {
	servic services.ServicesAutentikasi
}

func NewControlAuntentikasi(servic services.ServicesAutentikasi) *ControlAuntentikasi {
	return &ControlAuntentikasi{servic: servic}
}

func (ctrl *ControlAuntentikasi) Registers(c *gin.Context) {
	type Inputan struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var input Inputan
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Input gagal"})
		return
	}
	otp, err := ctrl.servic.Register(input.Email, input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Gagal"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": input.Email, "OTP": otp})
}

func (ctrl *ControlAuntentikasi) VerifysOtp(c *gin.Context) {
	var inputan struct {
		Email string `json:"email"`
		OTP   string `json:"otp"`
	}
	if err := c.ShouldBindJSON(&inputan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "inputan salah"})
		return
	}
	if err := ctrl.servic.VerifyOtp(inputan.Email, inputan.OTP); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "verifikasi gagal"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Berhasil Lanjut Login"})
}

func (ctrl *ControlAuntentikasi) LoginUser(c *gin.Context) {
	type input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var inputs input
	if err := c.ShouldBindJSON(&inputs); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "input gagal"})
		return
	}
	user, err := ctrl.servic.Login(inputs.Email, inputs.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Harus masukkan otp dulu"})
		return
	}
	token, err := middlewares.GenerateJwt(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Login Berhasil", "token": token})
}
