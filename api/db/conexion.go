// Package db establece una conexión con la base de datos
package db

import (
	_ "github.com/go-sql-driver/mysql" //
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //
)

//InitConection d
func InitConection() (db *gorm.DB, err error) {
	db, err = gorm.Open("mysql", "root:1524@(localhost:3308)/soporte_tickets?charset=utf8&parseTime=True&loc=Local")
	return
}
