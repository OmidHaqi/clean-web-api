package migrations

import (
	"github.com/omidhaqi/clean-web-api/config"
	"github.com/omidhaqi/clean-web-api/constants"
	"github.com/omidhaqi/clean-web-api/data/models"
	"github.com/omidhaqi/clean-web-api/infra/persistence/database"
	"github.com/omidhaqi/clean-web-api/pkg/logging"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var logger = logging.NewLogger(config.GetConfig())

func Up_1() {
	database := database.GetDb()

	createTables(database)
	createDefaultInformation(database)

}

func createTables(database *gorm.DB) {

	tables := []interface{}{}

	country := models.Country{}
	city := models.City{}
	user := models.User{}
	role := models.Role{}
	userRole := models.UserRole{}

	tables = addNewTable(database, country, tables)
	tables = addNewTable(database, city, tables)
	tables = addNewTable(database, user, tables)
	tables = addNewTable(database, role, tables)
	tables = addNewTable(database, userRole, tables)

	database.Migrator().CreateTable(tables...)
	logger.Info(logging.Postgres, logging.Migration, "", nil)
}

func addNewTable(database *gorm.DB, models interface{}, tables []interface{}) []interface{} {
	if !database.Migrator().HasTable(models) {

		tables = append(tables, models)

	}
	return tables
}

func createDefaultInformation(database *gorm.DB) {


	adminRole := models.Role{Name: constants.AdminRoleName}
	createRoleIfNotExits(database, &adminRole)
	defaultRole := models.Role{Name: constants.DefaultRoleName}
	createRoleIfNotExits(database, &defaultRole)
	u := models.User{Username: constants.DefaultRoleUserName,FirstName: "Umut",LastName: "Haqi",MobileNumber: "09100090780",Email: "omid.haqi@outlook.com"}

	pass := "12345678"
	hashPass , _ := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	u.Password = string(hashPass)
	createAdminUserIfNotExits(database,&u,adminRole.Id)

}

func createRoleIfNotExits(database *gorm.DB, r *models.Role) {
	exists := 0
	database.
		Model(&models.Role{}).
		Select("1").
		Where("username = ?", r.Name).
		First(&exists)
	if exists == 0 {
		database.Create(r)
	}
}

func createAdminUserIfNotExits(database *gorm.DB, u *models.User,roleId int) {
	exists := 0
	database.
		Model(&models.User{}).
		Select("1").
		Where("name = ?", u.Username).
		First(&exists)
	if exists == 0 {
		database.Create(u)
		ur:= models.UserRole{UserId: u.Id,RoleId: roleId}
		database.Create(ur)
	}
}

func Down_1() {

}
