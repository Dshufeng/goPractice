package main

import (
	"fmt"

	"goPractice/model"

	"crypto/md5"
	"encoding/hex"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	//db := conn()
	//defer db.Close()
	//
	//var users []model.User
	//fmt.Println(db)
	//db.Table("zzx_user").Last(&users)
	//
	//fmt.Println(users)
	//login()
	Edit()
}

func conn() (db *gorm.DB) {
	args := "wangjk:wangjk_21@web@tcp(47.92.23.244:3306)/zzx_demo?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", args)

	if err != nil {
		panic("failed to connect database")
	}
	return
}

func login() {
	var name string
	var password string
	fmt.Print("username: ")
	fmt.Scanln(&name)

	fmt.Print("password: ")
	fmt.Scan(&password)

	db := conn()
	db.LogMode(true)
	defer db.Close()

	var people model.People
	err := db.Table("zzx_user").Where("username = ?", name).First(&people).Error
	if err == gorm.ErrRecordNotFound {
		fmt.Println("用户名不存在")
		return
	}

	if (MD5HEX(MD5HEX(password) + people.Salt)) != people.Password {
		fmt.Println("login err")
		return
	}
	fmt.Println("login success")
}

func MD5HEX(str string) string {
	s := md5.New()
	s.Write([]byte(str))
	return hex.EncodeToString(s.Sum(nil))
}

func Edit() {
	db := conn()
	people := &model.People{
		UID: 1,
	}
	db.Table("zzx_user").Model(&people).Update("reg_ip", "127.1.1.1")

}
