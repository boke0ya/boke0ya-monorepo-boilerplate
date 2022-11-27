package repositories

import (
	"app/internal/entities"
	"app/internal/errors"
	"app/internal/infrastructures/gorm/dao"
	"gorm.io/gorm"
)

type GormEmailVerificationRepository struct {
	db *gorm.DB
}

func NewGormEmailVerificationRepository(db *gorm.DB) entities.EmailVerificationRepository {
	return GormEmailVerificationRepository{db}
}

func (cr GormEmailVerificationRepository) Begin() entities.EmailVerificationRepository {
	return GormEmailVerificationRepository{
		db: cr.db.Begin(),
	}
}

func (cr GormEmailVerificationRepository) Commit() error {
	if err := cr.db.Commit().Error; err != nil {
		return errors.New(errors.FailedToPersistEmailVerification, err)
	}
	return nil
}

func (cr GormEmailVerificationRepository) Rollback() {
	cr.db.Rollback()
}

func (cr GormEmailVerificationRepository) Create(user entities.EmailVerification) (entities.EmailVerification, error) {
	gormEmailVerification := dao.ConvertFromEmailVerificationEntity(user)
	if err := cr.db.Create(&gormEmailVerification).Error; err != nil {
		return entities.EmailVerification{}, errors.New(errors.FailedToPersistEmailVerification, err)
	}
	return dao.ConvertToEmailVerificationEntity(gormEmailVerification), nil
}

func (cr GormEmailVerificationRepository) Update(user entities.EmailVerification) (entities.EmailVerification, error) {
	gormEmailVerification := dao.ConvertFromEmailVerificationEntity(user)
	if err := cr.db.Save(&gormEmailVerification).Error; err != nil {
		return entities.EmailVerification{}, errors.New(errors.FailedToPersistEmailVerification, err)
	}
	return dao.ConvertToEmailVerificationEntity(gormEmailVerification), nil
}

func (cr GormEmailVerificationRepository) DeleteById(id string) error {
	if err := cr.db.Delete(&dao.EmailVerification{}, "id = ?", id).Error; err != nil {
		return errors.New(errors.FailedToPersistEmailVerification, err)
	}
	return nil
}

func (cr GormEmailVerificationRepository) FindById(id string) (entities.EmailVerification, error) {
	var user dao.EmailVerification
	if err := cr.db.First(&user, "id = ?", id).Error; err != nil {
		return entities.EmailVerification{}, errors.New(errors.EmailVerificationNotFoundError, err)
	}
	return dao.ConvertToEmailVerificationEntity(user), nil
}

func (cr GormEmailVerificationRepository) FindByToken(token string) (entities.EmailVerification, error) {
	var user dao.EmailVerification
	if err := cr.db.First(&user, "token = ?", token).Error; err != nil {
		return entities.EmailVerification{}, errors.New(errors.EmailVerificationNotFoundError, err)
	}
	return dao.ConvertToEmailVerificationEntity(user), nil
}
