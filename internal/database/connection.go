package database

import (
	"fmt"

	"awesome-content/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDbClient(config *config.AppConfig) (db *gorm.DB, err error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		config.DB.Username, config.DB.Password, config.DB.Host, config.DB.Port,
		config.DB.Name, config.DB.Charset,
	)

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: false,
		PrepareStmt:            true,
	})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func CloseDbClient() {
	connection, _ := db.DB()
	err := connection.Close()
	if err != nil {
		panic(err)
	}
}
