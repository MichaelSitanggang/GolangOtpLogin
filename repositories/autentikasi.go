package repositories

import (
	"latihanotp/models"

	"gorm.io/gorm"
)

type RepoAutentikasi interface {
	FindByEmail(email string) (*models.User, error)
	Save(user *models.User) error
	CreateUser(user *models.User) error
}

type repoAutentikasi struct {
	db *gorm.DB
}

func NewRepoAutentikasi(db *gorm.DB) RepoAutentikasi {
	return &repoAutentikasi{db: db}
}

func (r *repoAutentikasi) FindByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, nil
	}
	return &user, nil
}

func (r *repoAutentikasi) Save(user *models.User) error {
	return r.db.Save(user).Error
}

func (r *repoAutentikasi) CreateUser(user *models.User) error {
	return r.db.Create(user).Error
}
