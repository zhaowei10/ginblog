package model

import (
	"encoding/base64"
	"ginblog/utils/errmsg"
	"golang.org/x/crypto/scrypt"
	"gorm.io/gorm"
	"log"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username"`
	Password string `gorm:"type:varchar(20);not null" json:"password"`
	Role     int    `gorm:"type:int" json:"role"`
}

//查询用户是否存在
func CheckUser(name string) (code int) {
	var user User
	db.Select("id").Where("username = ?", name).First(&user)
	if user.ID > 0 {
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCESS
}

//新增用户
func CreateUser(user *User) int {
	//user.Password = ScryptPwd(user.Password) //对密码加密
	err := db.Create(user).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//查询用户列表
func GetUserList(pageSize, pageNum int) ([]User, int64) {
	var userlist []User
	var total int64
	err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&userlist).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	return userlist, total
}

//编辑用户
func EditUser(id int, user *User) int {
	var user1 User
	var maps = make(map[string]interface{})
	maps["username"] = user.Username
	maps["role"] = user.Role
	err := db.Model(&user1).Where("id=?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//删除用户
func DeleteUser(id int) int {
	var user User
	db.Where("id=?", id).Find(&user)
	err := db.Delete(&user).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//密码加密
func (user *User) BeforeSave(_ *gorm.DB) (err error) {
	user.Password = ScryptPwd(user.Password)
	return nil
}
func ScryptPwd(password string) string {
	const KeyLen = 10
	salt := make([]byte, 8)
	salt = []byte{12, 4, 6, 33, 6, 86, 89, 44}
	HashPwd, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, KeyLen)
	if err != nil {
		log.Fatal(err)
	}
	finalPwd := base64.StdEncoding.EncodeToString(HashPwd)
	return finalPwd
}

//登录验证
func CheckLogin(username, password string) int {
	var user User
	db.Where("username=?", username).First(&user)
	if user.ID == 0 {
		return errmsg.ERROR_USER_NOT_EXIT
	}
	if ScryptPwd(password) != user.Password {
		return errmsg.ERROR_PASSWORD_WRONG
	}
	if user.Role != 1 {
		return errmsg.ERROR_USER_NORIGHT
	}
	return errmsg.SUCCESS
}
