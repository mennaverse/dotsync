package middleware

import (
	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v5"
	"github.com/mennaverse/dotsync/apps/backend/manager"
	"github.com/mennaverse/dotsync/apps/backend/service"
)

func Services() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			secrets := manager.NewSecretManager()
			dsn, err := secrets.GetDatabaseDSN()
			if err != nil {
				return err
			}

			ctx := c.Request().Context()
			db, err := pgx.Connect(ctx, dsn)
			if err != nil {
				return err
			}
			defer db.Close(ctx)

			tx, err := db.Begin(ctx)
			if err != nil {
				return err
			}
			defer tx.Rollback(ctx)

			services := service.NewServices(tx)
			c.Set("services", services)

			err = next(c)
			if err != nil {
				return err
			}

			err = tx.Commit(ctx)
			if err != nil {
				return err
			}

			return nil
		}
	}
}
