package services

import (
	"errors"
	"fmt"
	"latihanotp/models"
	"latihanotp/repositories"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type ServicesAutentikasi interface {
	SendOtpGmail(email string) (string, error)
	Register(email, password string) (string, error)
	VerifyOtp(email, otp string) error
	Login(email, password string) (*models.User, error)
}

type servicesAutentikasi struct {
	repo repositories.RepoAutentikasi
}

func NewServicesAuntentikasi(repo repositories.RepoAutentikasi) ServicesAutentikasi {
	return &servicesAutentikasi{repo: repo}
}

func (s *servicesAutentikasi) SendOtpGmail(email string) (string, error) {
	rand.Seed(time.Now().UnixNano())
	otp := rand.Intn(999999-100000) + 100000
	otpStr := strconv.Itoa(otp)

	user, _ := s.repo.FindByEmail(email)
	user.OTP = otpStr
	if err := s.repo.Save(user); err != nil {
		return "", err
	}

	from := mail.NewEmail("Karbon", "michaelsitanggang37@gmail.com")
	subject := "Otp Gmail Karbon"
	to := mail.NewEmail("User", email)
	plaintextcontext := " Otp Kamu " + otpStr
	htmlContent := "<strong>Otp kamu adalah " + otpStr + "</strong>"
	message := mail.NewSingleEmail(from, subject, to, plaintextcontext, htmlContent)
	ApiClient := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := ApiClient.Send(message)
	if err != nil {
		return "", err
	}
	if response.StatusCode >= 200 && response.StatusCode < 300 {
		return otpStr, nil
	}
	return "", fmt.Errorf("data tidak ada")
}

func (s *servicesAutentikasi) Register(email, password string) (string, error) {
	cekEmail, err := s.repo.FindByEmail(email)
	if err != nil {
		return "", err
	}
	if cekEmail != nil {
		return "", errors.New("email sudah ada")
	}
	user := &models.User{
		Email:    email,
		Password: password,
	}
	if err := s.repo.CreateUser(user); err != nil {
		return "", err
	}

	otp, err := s.SendOtpGmail(user.Email)
	if err != nil {
		return "", err
	}
	user.OTP = otp
	if err := s.repo.Save(user); err != nil {
		return "", nil
	}
	return otp, nil
}

func (s *servicesAutentikasi) VerifyOtp(email, otp string) error {
	user, _ := s.repo.FindByEmail(email)
	if otp == "" || otp != user.OTP {
		return fmt.Errorf("otp salah")
	}
	user.IsVerified = true
	if err := s.repo.Save(user); err != nil {
		return err
	}
	return nil
}

func (s *servicesAutentikasi) Login(email, password string) (*models.User, error) {
	user, _ := s.repo.FindByEmail(email)
	if user.Password != password {
		return nil, fmt.Errorf("password salah")
	}
	if !user.IsVerified {
		return nil, fmt.Errorf("akun belum terverifikasi")
	}
	return user, nil
}
