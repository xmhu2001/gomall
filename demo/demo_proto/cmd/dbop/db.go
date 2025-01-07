package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/xmhu2001/gomall/demo/demo_proto/biz/dal"
	"github.com/xmhu2001/gomall/demo/demo_proto/biz/dal/mysql"
	"github.com/xmhu2001/gomall/demo/demo_proto/biz/model"
)

func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		panic("Error loading .env file")
	}
	// 数据库访问层初始化
	dal.Init()
	// CURD
	// 数据创建：create单条数据；批量插入
	mysql.DB.Create(&model.User{
		Email:    "test@test.com",
		Password: "test_demo",
	})
	// 数据修改：修改单列；修改多列
	mysql.DB.Model(&model.User{}).Where("email = ?", "test@test.com").Update("password", "helloworld")
	// 读取
	row := model.User{}
	mysql.DB.Model(&model.User{}).Where("email = ?", "test@test.com").First(&row)
	fmt.Printf("row: %+v\n", row)

	// 删除
	mysql.DB.Unscoped().Where("email = ?", "test@test.com").Delete(&model.User{})
}
