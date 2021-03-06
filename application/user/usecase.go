package user

import (
	"github.com/go-park-mail-ru/2020_1_k-on/application/models"
)

type UseCase interface {
	Add(usr *models.User) (*models.User, error)
	Get(id uint) (*models.User, error)
	Update(user *models.User) error
	SetImage(id uint, image string) error
	Oauth(vkUser *models.VkUser) (*models.User, error)
	GetOauthConfig() (*models.OauthConfig, error)
}
