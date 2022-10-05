package dao

import (
	"TaskSubmit/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type manager struct {
	db *gorm.DB
}

var Manager interface {
	Register(user *model.User)
	GetUsername(username string) model.User
	GetUserCookie(cookie string) model.User
}

// 初始化，链接数据库
func init() {
	// 设置数据库dsn，
	dsn := "root:zks123456789.Z@tcp(localhost:3306)/userinfo1?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to init db:", err)
	}
	Manager = &manager{db: db}
	db.AutoMigrate(&model.User{})
}

func (mgr *manager) Register(user *model.User) {
	mgr.db.Create(user)
}

func (mgr *manager) GetUsername(username string) model.User {
	var user model.User
	mgr.db.Where("username=?", username).First(&user)
	return user
}

func (mgr *manager) GetUserCookie(cookie string) model.User {
	var user model.User
	mgr.db.Where("cookie=?", cookie).First(&user)
	return user
}
