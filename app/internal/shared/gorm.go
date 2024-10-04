package shared

import (
	"fmt"
	"github.com/Nukie90/my-fluffy/app/domain/entity"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/spf13/viper"
)

type GormConfig struct {
	Driver string
	Host   string
	Port   string
	User   string
	Pass   string
	DBName string
}

func NewGormConfig(v *viper.Viper) *GormConfig {
	return &GormConfig{
		Driver: v.GetString("db.driver"),
		Host:   v.GetString("db.host"),
		Port:   v.GetString("db.port"),
		User:   v.GetString("db.user"),
		Pass:   v.GetString("db.pass"),
		DBName: v.GetString("db.dbname"),
	}
}

func (gc *GormConfig) Connector() (*gorm.DB, error) {
	var dsn string
	switch gc.Driver {
	case "sqlite3":
		dsn = fmt.Sprintf("%s.db", "./db/"+gc.DBName)
		db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
		if err != nil {
			return nil, err
		}
		gc.AutoMirage(db)

		return db, nil
	default:
		return nil, fmt.Errorf("unsupported driver: %s", gc.Driver)
	}
}

func (gc *GormConfig) AutoMirage(db *gorm.DB) {
	db.AutoMigrate(&entity.User{})
	db.AutoMigrate(&entity.Notification{})
	db.AutoMigrate(&entity.Post{})
}
