package mysql

import (
	"fmt"
	"os"

	"github.com/xmhu2001/gomall/app/product/biz/model"
	"github.com/xmhu2001/gomall/app/product/conf"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	dsn := fmt.Sprintf(conf.GetConf().MySQL.DSN, os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"))
	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}
	if os.Getenv("GO_ENV") != "online" {
		needDemoData := !DB.Migrator().HasTable(&model.Product{})
		DB.AutoMigrate( //nolint:errcheck
			&model.Product{},
			&model.Category{},
		)
		if needDemoData {
			DB.Exec("INSERT INTO `product`.`category` VALUES (1,'2023-12-06 15:05:06','2023-12-06 15:05:06','manga','manga'),(2,'2023-12-06 15:05:06','2023-12-06 15:05:06','shouban','shouban')")
			DB.Exec("INSERT INTO `product`.`product` VALUES ( 1, '2023-12-06 15:26:19', '2023-12-09 22:29:10', 'manga', 'test1', '/static/image/manga.png', 9.90 ), ( 2, '2023-12-06 15:26:19', '2023-12-09 22:29:10', 'shouban', 'test2', '/static/image/zhoubian.png', 6.60 )")
			DB.Exec("INSERT INTO `product`.`product_category` (product_id,category_id) VALUES ( 1, 1 ), ( 2, 2 )")
		}
	}
}
