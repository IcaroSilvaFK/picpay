package factories

import "github.com/IcaroSilvaFK/picpay/infra/database"

func NewDbConfig() database.Config {
	return database.Config{
		Provider: "postgres",
		DSN:      "host=localhost user=docker password=docker dbname=picpay port=5432 sslmode=disable",
	}
}
