package errorcodehandling

import (
	"github.com/retere/IOTSmartMeasureKWH/config/app"
	"os"
)

type CodeError struct{}

func (*CodeError) ParseSQLError(err error) error {
	driver := os.Getenv("DB_DRIVER")

	switch driver {
	case "postgres":
		return app.ParsePostgresSQLError(err)
	}
	return err
}
