package repository

import (
	"github.com/go-park-mail-ru/2020_1_k-on/internal/models"
	"github.com/go-park-mail-ru/2020_1_k-on/internal/user"
	"github.com/jinzhu/gorm"
)

type UserDatabase struct {
	conn *gorm.DB
}

func NewUserDatabase(db *gorm.DB) user.Repository {
	return &UserDatabase{conn: db}
}

func (udb *UserDatabase) Add(user *models.User) (err error) {
	return udb.conn.Table("kinopoisk.users").Create(user).Error
}

func (udb *UserDatabase) Update(id int64, upUser *models.User) error {
	usr := new(models.User)
	udb.conn.Table("kinopoisk.users").Where("id = ?", id).First(usr)
	if upUser.Username != "" {
		usr.Username = upUser.Username
	}
	if upUser.Password != "" {
		usr.Password = upUser.Password
	}
	if upUser.Email != "" {
		usr.Email = upUser.Email
	}
	return udb.conn.Table("kinopoisk.users").Save(usr).Error
}

func (udb *UserDatabase) GetById(id int64) (usr *models.User, err error) {
	usr = new(models.User)
	err = udb.conn.Table("kinopoisk.users").Where("id = ?", id).First(usr).Error
	return usr, err
}

func (udb *UserDatabase) GetByName(login string) (usr *models.User, err error) {
	usr = new(models.User)
	err = udb.conn.Table("kinopoisk.users").Where("username = ?", login).First(usr).Error
	return usr, err
}

func (udb *UserDatabase) Contains(login string) (bool, error) {
	_, err := udb.GetByName(login)

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}
