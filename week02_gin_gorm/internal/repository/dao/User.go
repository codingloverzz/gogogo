package dao

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

type UserDao struct {
	db *gorm.DB
}

type User struct {
	Id       int64  `gorm:"primary_key autoIncrement"`
	Email    string `gorm:"unique"`
	AboutMe  string
	Password string
	//创建时间
	CTime int64
	//更新时间
	UTime int64
}

var (
	EmailDuplicateError = errors.New("邮箱冲突了啊")
	RecordNotFoundError = gorm.ErrRecordNotFound
)

func (dao *UserDao) Insert(ctx context.Context, user *User) error {
	fmt.Println("走进来了啊阿啊")

	now := time.Now().UnixMilli()
	user.CTime = now
	user.UTime = now
	err := dao.db.WithContext(ctx).Create(&user).Error

	var mysqlError *mysql.MySQLError
	if errors.As(err, &mysqlError) {
		const uniqueErrorNumber = 1062
		if mysqlError.Number == uniqueErrorNumber {
			return EmailDuplicateError
		}
	}
	return err

}

func (dao *UserDao) FindByEmail(ctx context.Context, email string) (User, error) {
	var u User
	err := dao.db.WithContext(ctx).Where("email=?", email).First(&u).Error
	return u, err
}

func (dao *UserDao) UpdateById(ctx context.Context, u *User) error {
	err := dao.db.WithContext(ctx).Where("id=?", u.Id).Updates(&User{
		AboutMe: u.AboutMe,
	}).Error
	return err
}

func NewUserDao(db *gorm.DB) *UserDao {
	return &UserDao{
		db: db,
	}
}
