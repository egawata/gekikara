package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

var (
	Db *gorm.DB
)

func InitDB() {
	var err error
	Db, err = gorm.Open("mysql", "root:@/gekikara?parseTime=true&loc=Asia%2FTokyo")
	if err != nil {
		panic(err)
	}

	Db.AutoMigrate(&Shop{}, &User{})
	log.Println("shop created")
}
