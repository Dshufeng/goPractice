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
	login()
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
	defer db.Close()

	var user model.User
	db.Table("zzx_user").Where("username = ?", name).First(&user)

	if (MD5HEX(MD5HEX(password) + user.Salt)) != user.Password {
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
