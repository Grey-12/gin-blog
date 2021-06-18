package models

import (
	_ "fmt"
	"github.com/Grey-12/gin-blog/pkg/setting"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"time"
)

var db *gorm.DB

type Model struct {
	ID uint `gorm:"primaryKey;autoIncrement;not null" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

func init() {
	var (
		err                                               error
		dbName, user, password, tablePrefix string
		//port 											  int
	)
	sec, err := setting.Cfg.GetSection("postgresql")
	if err != nil {
		log.Fatalln(2, "Fail to get section 'mysql': %v", err)
	}
	// dbType = sec.Key("TYPE").String()
	dbName = sec.Key("DBNAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	//host = sec.Key("HOST").String()
	//port, _ = sec.Key("PORT").Int()
	tablePrefix = sec.Key("TABLE_PREFIX").String()
	// dsn := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable", host, port, user, dbName, password)
	dsn := "host=localhost user=" + user + " password=" + password + " dbname=" + dbName + " port=5432  sslmode=disable TimeZone=Asia/Shanghai"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: tablePrefix,
			SingularTable: true,
		},
	})
	if err != nil {
		log.Println(err)
	}
	//sqlDB, err := db.DB()
	//sqlDB.SetMaxIdleConns(10)
	//sqlDB.SetMaxOpenConns(100)

	InitModels()
}

func InitModels() {
	err := db.AutoMigrate(&Tag{}, &Article{})
	if err != nil {
		log.Printf("数据库迁移出错 err: %v", err)
	}

}