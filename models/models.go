package models

import (
	"fmt"
	"github.com/Grey-12/gin-blog/pkg/setting"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
)

var db *gorm.DB

type Model struct {
	ID int `gorm:"primary_key" json:"id"`
	CreateOn int `json:"create_on"`
	ModifiedOn int `json:"modified_on"`
	DeletedOn  int `json:"deleted_on"`
}

func init() {
	var (
		err                                               error
		dbName, user, password, host, tablePrefix string
		port 											  int
	)
	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatalln(2, "Fail to get section 'mysql': %v", err)
	}
	// dbType = sec.Key("TYPE").String()
	dbName = sec.Key("NAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()
	port, _ = sec.Key("PORT").Int()
	tablePrefix = sec.Key("TABLE_PREFIX").String()
	dsn := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable", host, port, user, dbName, password)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: tablePrefix,
			SingularTable: true,
		},
	})
	if err != nil {
		log.Println(err)
	}
	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	// InitModels()
}

func InitModels() {
}