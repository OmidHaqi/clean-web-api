package migrations

import (
	"github.com/omidhaqi/clean-web-api/config"
	"github.com/omidhaqi/clean-web-api/data/models"
	"github.com/omidhaqi/clean-web-api/infra/persistence/database"
	"github.com/omidhaqi/clean-web-api/pkg/logging"
)

var logger = logging.NewLogger(config.GetConfig()) 

func Up_1() {

	database := database.GetDb()

	tables := []interface{}{}

	country := models.Country{}
	city := models.City{}

	if !database.Migrator().HasTable(country) {

		tables = append(tables, country)

	}
	if !database.Migrator().HasTable(city) {

		tables = append(tables, city)

	}

	database.Migrator().CreateTable(tables...)
	logger.Info(logging.Postgres,logging.Migration,"",nil)

}

func Down_1() {

}
