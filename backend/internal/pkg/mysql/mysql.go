package mysql

import (
	"fmt"
	"log"
	"presolai/configs"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(config *configs.Mysql) *gorm.DB {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.DataBase,
	)

	log.Println(dsn)

	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:               dsn,
		DefaultStringSize: 256, // default size for string fields
	}), &gorm.Config{})

	if err != nil {
		panic(`😫: Connected failed, check your Mysql with ` + dsn)
	}

	// Migrate the schema
	// migrateErr := db.AutoMigrate(&models.User{})
	// if migrateErr != nil {
	// 	panic(`😫: Auto migrate failed, check your Mysql with ` + address)
	// }

	// export DB
	DB = db
	log.Printf(`🍟 Successfully connected to Mysql`)

	return db

}
