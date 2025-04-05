package configs

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DBTOPAZ *gorm.DB

func InitDBTopaz() {
	dsn := os.Getenv("DBTOPAZDSN")
	var err error
	DBTOPAZ, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
}
