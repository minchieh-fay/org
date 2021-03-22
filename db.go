package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBEngine struct {
	// ID        uint `gorm:"primarykey"`
	name     string
	addr     string
	user     string
	password string
	gdb      *gorm.DB
}

var dbe *DBEngine

func (d *DBEngine) setDsn(name, addr, user, password string) {
	d.addr = addr
	d.name = name
	d.password = password
	d.user = user
}

func (d *DBEngine) connect() error {
	strDsn := fmt.Sprintf("%s:%s@tcp(%s)/?charset=utf8mb4&parseTime=True&loc=Local", d.user, d.password, d.addr)
	db, err := gorm.Open(mysql.Open(strDsn), &gorm.Config{})
	if err != nil {
		return err
	}
	sql := fmt.Sprintf("Create Database If Not Exists %s Character Set UTF8", d.name)
	db.Exec(sql)
	return nil
}

func (d *DBEngine) initTables() error {
	strDsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", d.user, d.password, d.addr, d.name)
	db, err := gorm.Open(mysql.Open(strDsn), &gorm.Config{})
	if err != nil {
		return err
	}

	db.AutoMigrate(&User{})

	var user = User{}
	user.UUID = newUUID()
	user.Name = "admin"
	user.LoginID = "admin"
	user.Password = "1:admin123"

	db.Where(User{LoginID: "admin"}).FirstOrCreate(&user)

	d.gdb = db
	return nil
}

/*******************************************/
func (d *DBEngine) getUserByLoginID(loginid string) *User {
	var user User
	tx := d.gdb.Where(User{LoginID: loginid}).First(&user)

	if tx.Error != nil {
		return nil
	}

	return &user
}
