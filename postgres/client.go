package postgres

import (
	"fmt"

	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Client struct {
	DB       *gorm.DB
	Host     string
	Port     int
	User     string
	Password string
	DBname   string
	Timezone string //Asia/Shanghai
}

// Options is a Client configuration function
type Options func(*Client)

// NewWorkClient return a ready to use client instance.
func NewClient(opts ...Options) *Client {
	c := &Client{
		Host:     "localhost",
		Port:     5432,
		User:     "postgres",
		Password: "postgres",
		DBname:   "buildeploydb",
		Timezone: "Europe/Paris",
	}

	// Apply configuration
	for _, opt := range opts {
		opt(c)
	}

	return c
}

func (c *Client) Open() error {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=%s",
		c.Host, c.Port, c.User, c.Password, c.DBname, c.Timezone)
	//dsn := "user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})

	if err != nil {
		return errors.Wrap(err, "gorm DB open failed")
	}

	//db.LogMode(true)

	c.DB = db

	return nil
}

func (c *Client) AutoMigrate(object interface{}) error {
	if err := c.DB.AutoMigrate(object).Error; err != nil {
		return fmt.Errorf(err())
	}

	return nil
}
