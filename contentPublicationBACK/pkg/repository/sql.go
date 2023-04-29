package repository

import (
	app "contentPublicationBACK"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	titleTableE        = "Titles"
	categoryTableE     = "Categories"
	tagTableE          = "Tags"
	serialTableE       = "Serials"
	titleContentTableE = "Content"
	imagesTableE       = "Content.Images"
	likeTableE         = "Content.Likes"
	userTableE         = "Content.Likes.User"
	commentTableE      = "Content.Comments"
	commentUserTableE  = "Content.Comments.User"
)

type Config struct {
	Host       string
	DriverName string
	Port       string
	DBName     string
	Username   string
	Password   string
	ParseTime  string
}

func NewDB(cfg Config) (*gorm.DB, error) {
	db, err := connectToDb(cfg)
	return db, err
}

func CreateORMModels(config Config) {
	db, _ := connectToDb(config)
	db.AutoMigrate(app.User{})
	db.AutoMigrate(app.Title{})
	db.AutoMigrate(app.Category{})
	db.AutoMigrate(app.Tag{})
	db.AutoMigrate(app.TitleContent{})
	db.AutoMigrate(app.Image{})
	db.AutoMigrate(app.Like{})
	db.AutoMigrate(app.StaticType{})
	db.AutoMigrate(app.Comment{})
	db.AutoMigrate(app.Subscription{})
}

func connectToDb(cfg Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=%s&loc=Local",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.ParseTime)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	gormDb, err := db.DB()

	err = gormDb.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
