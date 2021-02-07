package models

import (
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	db  *gorm.DB
	err error
)

//NewDatabase 数据库连接初始化器
func NewDatabase() error {
	driverName := viper.GetString("database.type")
	user := viper.GetString("database.user")
	password := viper.GetString("database.password")
	port := viper.GetString("database.port")
	dbname := viper.GetString("database.dbname")
	config := viper.GetString("database.config")
	dsn := user + ":" + password + "@tcp(:" + port + ")/" + dbname + "?" + config

	db, err = gorm.Open(mysql.New(mysql.Config{
		DriverName: driverName,
		DSN:        dsn,
		//DefaultStringSize:         256,   // string 类型字段的默认长度
		//DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		//DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		//DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		//SkipDefaultTransaction: true, //默认事务，禁用可以提升30%性能
		//DisableForeignKeyConstraintWhenMigrating: true,  //禁用gorm默认的外键约束
		NamingStrategy: schema.NamingStrategy{ ////命名策略表、列的命名策略
			SingularTable: true, //禁用数据库表名复数
			//TablePrefix:   "blog_", //表名前缀
		},
	})
	if err != nil {
		logrus.WithField("database", driverName).Error("Failed to open MySQL!")
		return err
	}

	//数据库连接池
	sqlDB, err := db.DB()
	if err != nil {
		logrus.WithField("err", err).Error("Failed to get database connection pool!")
		return err
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	// //数据库迁移
	err = db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4").AutoMigrate(&User{}, &Article{}, &Tag{})
	if err != nil {
		logrus.WithField("err", err).Error("Failed to migrate database!")
		return err
	}
	return nil
}
