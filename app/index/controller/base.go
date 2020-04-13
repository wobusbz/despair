package controller

import "despair/common/jwt"

type BaseController struct {
	Auth *jwt.AuthToken
}

func NewBaseController() *BaseController {
	return &BaseController{Auth: jwt.NewAuthToken()}
}
