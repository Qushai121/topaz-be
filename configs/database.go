package configs

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type BASEDB struct {
	TOPAZDB *gorm.DB
}

var DB BASEDB

func InitDBTopaz() {
	dsn := os.Getenv("DBTOPAZDSN")
	var err error
	DB.TOPAZDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
}
