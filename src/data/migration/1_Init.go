package migration

import (
	"github.com/mhmojtaba/golang-car-web-api/config"
	"github.com/mhmojtaba/golang-car-web-api/constants"
	"github.com/mhmojtaba/golang-car-web-api/data/db"
	"github.com/mhmojtaba/golang-car-web-api/data/models"
	"github.com/mhmojtaba/golang-car-web-api/pkg/logging"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var logger = logging.NewLogger(config.GetConfig())

func Up_1() {

	dataBase := db.GetDb()
	createTables(dataBase)
	createDefaultInfo(dataBase)

}

func createTables(dataBase *gorm.DB) {
	tables := []interface{}{}

	country := models.Country{}
	city := models.City{}
	user := models.User{}
	role := models.Role{}
	userRole := models.UserRole{}

	tables = addNewTable(dataBase, country, tables)
	tables = addNewTable(dataBase, city, tables)
	tables = addNewTable(dataBase, user, tables)
	tables = addNewTable(dataBase, role, tables)
	tables = addNewTable(dataBase, userRole, tables)

	dataBase.Migrator().CreateTable(tables...)
	logger.Info(logging.Postgres, logging.Migration, "tables created", nil)
}

func createDefaultInfo(dataBase *gorm.DB) {
	adminRole := models.Role{Name: constants.AdminRoleName}
	createRoleIfNotExist(dataBase, &adminRole)

	defaultRole := models.Role{Name: constants.AdminRoleName}
	createRoleIfNotExist(dataBase, &defaultRole)

	adminUser := models.User{Username: constants.DefaultUserName, FirstName: "Mojtaba", LastName: "mohammadi", Mobile: "09111111111", Email: "mmm@test.com"}
	pass := "admin123"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	adminUser.Password = string(hashedPassword)

	createAdminUserIfNotExist(dataBase, &adminUser, adminRole.Id)

}

func createRoleIfNotExist(database *gorm.DB, r *models.Role) {
	exists := 0
	database.
		Model(&models.Role{}).
		Select("1").
		Where("name = ?", r.Name).
		First(&exists)
	if exists == 0 {
		database.Create(r)
	}

}

func createAdminUserIfNotExist(database *gorm.DB, u *models.User, roleId int) {
	exists := 0
	database.
		Model(&models.User{}).
		Select("1").
		Where("username = ?", u.Username).
		First(&exists)
	if exists == 0 {
		database.Create(u)
		ur := models.UserRole{UserId: u.Id, RoleId: roleId}
		database.Create(&ur)
	}

}

func addNewTable(dataBase *gorm.DB, model interface{}, tables []interface{}) []interface{} {
	if !dataBase.Migrator().HasTable(model) {
		tables = append(tables, model)
	}
	return tables
}

func Down_1() {}
