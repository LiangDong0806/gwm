package service

import (
	"github.com/go-errors/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strconv"
	"zg5/Homework01/server/global"
)

var MysqlDB *gorm.DB
var err error

func InitMysql() {
	username := global.RpcALLConf.Mysql.Username
	password := global.RpcALLConf.Mysql.Password
	host := global.RpcALLConf.Mysql.Host
	port := strconv.Itoa(global.RpcALLConf.Mysql.Port)
	dbname := global.RpcALLConf.Mysql.Dbname
	dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"
	MysqlDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to open database")
	}
	MysqlDB.AutoMigrate(new(Product), new(User))
}

func PreheatTheProduct(name string) (*[]Product, error) {
	var products []Product
	err = MysqlDB.Where("name = ?", name).Find(&products).Error
	if err != nil {
		return nil, errors.New("数据查询失败")
	}
	return &products, nil
}

func ProductAddTo(pro Product) (Product, error) {
	res := MysqlDB.Create(&pro)
	return pro, res.Error
}

func QueryTheUser(username string) (User, error) { //查询用户
	var user User
	res := MysqlDB.Model(&User{}).Where("username = ?", username).First(&user)
	return user, res.Error
}

func UserRegistration(user User) (User, error) { //用户注册
	txt := MysqlDB.Begin()
	err = MysqlDB.Model(&User{}).Create(&user).Error
	if err != nil {
		txt.Rollback()
		return User{}, err
	}
	txt.Commit()
	return user, nil
}

type Product struct {
	Id          int    `gorm:"primaryKey"`
	Name        string `gorm:"not null;index"`
	Description string
	Price       float64 `gorm:"not null"`
	Stock       int     `gorm:"not null"`
	Category    string  // 商品类别
	gorm.Model
}

type Cart struct { //购物车表
	Id        int  `gorm:"primaryKey"`
	UserID    uint `gorm:"not null;index"` // 为用户ID字段添加索引
	ProductID uint `gorm:"not null;index"` // 为商品ID字段添加索引
	Quantity  int
	gorm.Model
}

type User struct {
	Id       int
	Username string
	Password string
	Mobile   string
	gorm.Model
}
