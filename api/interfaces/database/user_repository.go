package database

import (
	"github.com/watarun54/go-echo-gorm-clean-api/api/domain"
)

type UserRepository struct {
	SqlHandler
}

func (repo *UserRepository) FindById(id int) (user domain.User, err error) {
	if err = repo.Find(&user, id).Error; err != nil {
		return
	}
	return
}

func (repo *UserRepository) FindAll() (users domain.Users, err error) {
	if err = repo.Find(&users).Error; err != nil {
		return
	}
	return
}

func (repo *UserRepository) Store(u domain.User) (user domain.User, err error) {
	if err = repo.Create(&u).Error; err != nil {
		return
	}
	user = u
	return
}

func (repo *UserRepository) Update(u domain.User) (user domain.User, err error) {
	if err = repo.First(&domain.User{}, u.ID).Save(&u).Error; err != nil {
		return
	}
	user = u
	return
}

func (repo *UserRepository) DeleteById(user domain.User) (err error) {
	if err = repo.Delete(&user).Error; err != nil {
		return
	}
	return
}
