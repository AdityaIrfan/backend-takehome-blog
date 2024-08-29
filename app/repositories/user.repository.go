package repositories

import (
	contract "backend-takehome-blog/contracts"
	"backend-takehome-blog/models"
	"errors"
	"fmt"

	"github.com/phuslu/log"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) contract.IUserRepository {
	return &userRepository{
		db: db,
	}
}

func (u *userRepository) GetByCustomAndSelectedFields(values map[string]interface{}, selectedFields string) (*models.User, error) {
	db := u.db

	for field, value := range values {
		db = db.Where(fmt.Sprintf("%s = ?", field), value)
	}

	var user *models.User

	if err := db.Select(selectedFields).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		log.Error().Err(errors.New("FAILED TO GET USER : " + err.Error())).Msg("")
		return nil, err
	}

	return user, nil
}

func (u *userRepository) Create(user *models.User) error {
	if err := u.db.Clauses(clause.Returning{}).Create(&user).Error; err != nil {
		log.Error().Err(errors.New("FAILED TO CREATE USER : " + err.Error())).Msg("")
		return err
	}

	return nil
}
