package repositories

import (
	"github.com/boke0ya/beathub-api/internal/entities"
	"github.com/boke0ya/beathub-api/internal/errors"
	"github.com/boke0ya/beathub-api/internal/infrastructures/gorm/dao"
	"gorm.io/gorm"
)

type GormUserRepository struct {
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) entities.UserRepository {
	return GormUserRepository{db}
}


func (cr GormUserRepository) Create(user entities.User) (entities.User, error) {
	gormUser := dao.ConvertFromUserEntity(user)
	if err := cr.db.Create(&gormUser).Error; err != nil {
		return entities.User{}, errors.New(errors.FailedToPersistUser, err)
	}
	return dao.ConvertToUserEntity(gormUser), nil
}

func (cr GormUserRepository) Update(user entities.User) (entities.User, error) {
	gormUser := dao.ConvertFromUserEntity(user)
	if err := cr.db.Save(&gormUser).Error; err != nil {
		return entities.User{}, errors.New(errors.FailedToPersistUser, err)
	}
	return dao.ConvertToUserEntity(gormUser), nil
}

func (cr GormUserRepository) DeleteById(id string) error {
	if err := cr.db.Delete(&dao.User{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}

func (cr GormUserRepository) FindById(id string) (entities.User, error) {
	var user dao.User
	err := cr.db.First(&user, "id = ?", id).Error
	if err != nil {
		return entities.User{}, err
	}
	return dao.ConvertToUserEntity(user), err
}

func (cr GormUserRepository) FindByEmail(email string) (entities.User, error) {
	var user dao.User
	err := cr.db.First(&user, "email = ?", email).Error
	if err != nil {
		return entities.User{}, err
	}
	return dao.ConvertToUserEntity(user), err
}

