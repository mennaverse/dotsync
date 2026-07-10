package service

import (
	"github.com/jackc/pgx/v5"
	"github.com/mennaverse/dotsync/apps/backend/db"
	"github.com/mennaverse/dotsync/apps/backend/manager"
)

type Services struct {
	Authentication AuthenticationService
}

func NewServices(tx pgx.Tx) *Services {
	queries := db.New(tx)

	defaultManagers := manager.DefaultManagers()

	return &Services{
		Authentication: NewAuthenticationService(
			queries,
			defaultManagers.Secret,
			defaultManagers.Crypto,
			defaultManagers.Email,
		),
	}
}
