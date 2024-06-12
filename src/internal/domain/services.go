package domain

import (
	"github.com/rachitkawar/boilerplate-go/src/internal/database"
	"github.com/rachitkawar/boilerplate-go/src/internal/domain/auth"
)

type Service struct {
	Auth *auth.AuthSrv
}

func NewService(d database.Store) *Service {
	return &Service{
		Auth: auth.NewAuthSrv(d),
	}
}
