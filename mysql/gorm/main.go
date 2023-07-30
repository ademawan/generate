package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/lithammer/shortuuid"

	"github.com/labstack/gommon/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model `json:"-"`
	User_uid   string         `gorm:"index;unique;type:varchar(22)" json:"user_uid"`
	Name       string         `gorm:"type:varchar(100)" json:"name"`
	Email      string         `gorm:"unique" json:"email"`
	Image      string         `json:"image"`
	History    []User_history `gorm:"foreignKey:User_uid;references:User_uid" json:"history"`
}
type User_history struct {
	ID uint `gorm:"primarykey" json:"-"`

	User_history_uid string         `json:"user_history_uid"`
	User_uid         string         `gorm:"index;type:varchar(32)" json:"user_uid"`
	CreatedAt        time.Time      `json:"cretedAt"`
	UpdatedAt        time.Time      `json:"-"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"-"`
}

func main() {
	InitDB()

}

func InitDB() *gorm.DB {

	connectionString := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=%v",
		"root",
		"root",
		"localhost",
		"3308",
		"test_gorm",
		"Asia%2FJakarta",
	)
	fmt.Println(connectionString)
	DB, err := gorm.Open(mysql.Open(connectionString) /* &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true} */)

	if err != nil {
		log.Info("error in connect database : ", err)
		panic(err)
	}
	fmt.Println("connected")
	AutoMigrate(DB)
	return DB
}

func AutoMigrate(DB *gorm.DB) {
	DB.AutoMigrate(&User{})
	// 	DB.AutoMigrate(&Goal{})
	// 	DB.AutoMigrate(&Food{})
	// 	DB.AutoMigrate(&Menu{})
	// 	DB.AutoMigrate(&Detail_menu{})
	DB.AutoMigrate(&User_history{})
}
func GetById(db *gorm.DB, user_uid string) (User, error) {
	arrUser := User{}

	result := db.Preload("Goal").Preload("History").Where("user_uid =?", user_uid).First(&arrUser)
	if result.RowsAffected == 0 {
		return arrUser, errors.New("record not found")
	}
	if err := result.Error; err != nil {
		return arrUser, err
	}

	return arrUser, nil
}
func Register(db *gorm.DB, u User) (User, error) {

	uid := shortuuid.New()
	u.User_uid = uid

	if err := db.Create(&u).Error; err != nil {
		return u, errors.New("invalid input or this email was created (duplicated entry)")
	}

	return u, nil
}
func Insert(db *gorm.DB, newHistory User_history) (User_history, error) {

	err := db.Transaction(func(tx *gorm.DB) error {

		uid := shortuuid.New()
		newHistory.User_history_uid = uid

		if err := tx.Model(User_history{}).Create(&newHistory).Error; err != nil {
			return errors.New("")
		}

		return nil
	})

	if err != nil {
		return User_history{}, err
	}

	return newHistory, nil
}
