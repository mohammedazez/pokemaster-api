package mysql

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"pokemaster-api/utils/helper"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	//driver
	"gorm.io/driver/mysql"
)

var (
	//mysqlDB mysql db
	mysqlDB *gorm.DB
)

// Config mysql configuration
type Config struct {
	Host      string
	User      string
	Password  string
	DBName    string
	DBNumber  int
	Port      int
	DebugMode bool
}

// GetMysqlDB get mysql instance
func GetMysqlDB() (*gorm.DB, error) {
	if mysqlDB == nil {
		return nil, errors.New("please init mysql db")
	}

	return mysqlDB, nil
}

// LoadConfig load mysql configuration
func LoadConfig() Config {
	return Config{
		Host:      os.Getenv("DB_POKEMASTER_HOST"),
		User:      os.Getenv("DB_POKEMASTER_USERNAME"),
		Password:  os.Getenv("DB_POKEMASTER_PASSWORD"),
		DBName:    os.Getenv("DB_POKEMASTER_NAME"),
		Port:      helper.StringToInt(os.Getenv("DB_POKEMASTER_PORT"), 3306),
		DebugMode: helper.StringToBool(os.Getenv("DB_POKEMASTER_DEBUG"), false),
	}
}

// Init init mysql
func Init() {
	mysqlDB = MysqlConnect()
	sqlDB, err := mysqlDB.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(time.Hour)
}

// MysqlConnect connect to mysql using config name. return *gorm.DB incstance
func MysqlConnect() *gorm.DB {
	mysqlConfig := LoadConfig()

	dsn := fmt.Sprintf(`%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local`, mysqlConfig.User, mysqlConfig.Password, mysqlConfig.Host, strconv.Itoa(mysqlConfig.Port), mysqlConfig.DBName)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			IgnoreRecordNotFoundError: false,       // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,        // Disable color
		},
	)

	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		panic(err)
	}

	if mysqlConfig.DebugMode {
		return connection.Debug()
	}

	return connection
}
