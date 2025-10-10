package migration

import (
	"github.com/mhmojtaba/golang-car-web-api/config"
	"github.com/mhmojtaba/golang-car-web-api/data/db"
	"github.com/mhmojtaba/golang-car-web-api/data/models"
	"github.com/mhmojtaba/golang-car-web-api/pkg/logging"
	"github.com/mhmojtaba/golang-car-web-api/constants/constants"
)

var logger = logging.NewLogger(config.GetConfig())

func Up_1() {

	dataBase := db.GetDb()
	createTables(dataBase)
	createDefaultInfo(dataBase)
	
}

func createTables(dataBase *gorm.DB)  {
	tables := []interface{}{}

	country := models.Country{}
	city := models.City{}
	user := models.User{}
	role := models.Role{}
	userRole := models.UserRole{}

tables = addNewTable(dataBase,country,tables)
tables = addNewTable(dataBase,city,tables)
tables = addNewTable(dataBase,user,tables)
tables = addNewTable(dataBase,role,tables)
tables = addNewTable(dataBase,userRole,tables)

	dataBase.Migrator().CreateTable(tables...)
	logger.Info(logging.Postgres, logging.Migration, "tables created", nil)
}

func createDefaultInfo(dataBase *gorm.DB)  {
	exists := -1
	dataBase.Model(&models.Role{}).
		Select("1").
		Where("name=?",constants.AdminRoleName).
		First(&exists)

		if exists ==0 {
			r:= models.Role{Name: constants.AdminRoleName}
			dataBase.Create(&r)
		}else{
			return
		}
}

func addNewTable(dataBase *gorm.DB,model interface{},tables []interface{}){
			if !dataBase.Migrator().HasTable(model) {
		tables = append(tables, model)
	}
	return tables
}

func Down_1() {}
