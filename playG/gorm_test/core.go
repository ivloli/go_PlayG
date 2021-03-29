package main

import (
	"fmt"
	"time"

	"gopkg.in/guregu/null.v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type AmUser struct {
	UserName     string      `gorm:"column:UserName"`
	ProfilePhoto null.String `gorm:"column:ProfilePhoto"`
	UpdateTime   null.Time   `gorm:"column:UpdateTime"`
	Email        null.String `gorm:"column:Email"`
	UserType     int         `gorm:"column:UserType"`
	Team         string      `gorm:"column:Team"`
	ParentID     string      `gorm:"column:ParentId"`
	IsLeader     int         `gorm:"column:IsLeader"`
	RoleName     string      `gorm:"column:RoleName"`
	Credentials  null.String `gorm:"column:Credentials"`
	GroupName    string      `gorm:"column:GroupName"`
	Position     int         `gorm:"column:Position"`
	CreateTime   null.Time   `gorm:"column:CreateTime"`
	Phone        null.String `gorm:"column:Phone"`
	Path         string      `gorm:"column:Path"`
	ID           int         `gorm:"column:ID;primary_key"`
	UserID       string      `gorm:"column:UserId"`
	Status       null.Int    `gorm:"column:Status"`
}

// TableName sets the insert table name for this struct type
func (a *AmUser) TableName() string {
	return "am_user"
}

type LessUser struct {
	UserName string `gorm:"column:UserName"`
	UserType int    `gorm:"column:UserType"`
	RoleName string `gorm:"column:RoleName"`
	ID       int    `gorm:"column:ID;primary_key"`
	UserID   string `gorm:"column:UserId"`
	Section  string `gorm:"column:Section"`
}

func main() {
	db, err := gorm.Open(mysql.Open("adx:adx@tcp(127.0.0.1:3306)/am_db?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		fmt.Println("conn error")
		return
	}
	var results []AmUser
	db.Raw("select UserName,ParentID from am_user").Scan(&results)
	for _, v := range results {
		fmt.Println(v.UserName)
		fmt.Println(v.ParentID)
	}
	/*
		newUser := AmUser{}
		newUser.UserName = "abcfg"
		rownum := db.Model(&newUser).Where("UserName=?", newUser.UserName).Updates(AmUser{UserName: "oooooo", UserID: "sales101019"})
		fmt.Println(rownum)
		fmt.Println(newUser.UserName)
	*/
	var users []LessUser
	newUser := AmUser{}
	newUser.CreateTime = null.NewTime(time.Now(), true)
	newUser.RoleName = "IsNEW"
	newUser.Credentials = null.NewString("HAHAHA", true)
	newUser.UserName = "AAAAAB"
	newUser.UserID = "aaaa123457"
	result := db.Create(&newUser)
	fmt.Println(result)
	fmt.Println(newUser)
	db.Raw("select UserName,RoleName,ID,'abcd' as Section from am_user where ID > ? order by ID desc", 80).Find(&users)
	//fmt.Println(users)
	for _, v := range users {
		fmt.Println(v)
	}
	upUser := AmUser{
		UserID:   "abc123456zzz111",
		RoleName: "HackerJJJ",
		UserName: "xiaoji",
	}
	feilds := []string{"UserID", "RoleName"}
	where := &AmUser{
		UserName: "oooooo",
	}
	//db.Model(AmUser{}).Where(where).Select(feilds).Updates(&upUser)
	updatePackInfo(db, feilds, where, &upUser)
}

func updatePackInfo(db *gorm.DB, fields []string, filter, update schema.Tabler) error {
	res := db.Table(update.TableName()).Where(filter).Select(fields).Updates(update)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
