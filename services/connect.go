package services

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func init() {
	//open a db connection
	var err error
	DB, err = gorm.Open("mysql", "trungduc08:123456789@/cms_db?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Printf("err: %v", err)
		panic("failed to connect database")
	} else {
		fmt.Println("connect db sucessfully")
	}
}
