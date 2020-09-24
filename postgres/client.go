package postgres

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/souleb/buildeploy/app"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	ErrInvalidID = errors.New("gorm: ID provided was invalid")
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

// NewClient return a ready to use client instance.
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
	//dsn := "workflow=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return errors.Wrap(err, "gorm DB open failed")
	}

	//db.LogMode(true)

	c.DB = db

	return nil
}

// Read will read the provided workflow by ID.
func (c *Client) Read(id uint) (*app.Workflow, error) {
	var workflow app.Workflow
	err := c.DB.Where("id = ?", id).First(&workflow).Error
	if err != nil {
		return nil, errors.Wrap(err, "gorm: ID provided was invalid")
	}

	return &workflow, nil
}

// CreateWorkflow will create the provided workflow and backfill data
// like the ID, CreatedAt, and UpdatedAt fields.
func (c *Client) CreateWorkflow(workflow *app.Workflow) error {
	return c.DB.Create(workflow).Error
}

// Update will update the provided workflow with all of the data // in the provided workflow object.
func (c *Client) Update(workflow *app.Workflow) error {
	return c.DB.Save(workflow).Error
}

// Delete will delete the workflow with the provided ID
func (c *Client) Delete(id uint) error {
	if id == 0 {
		return fmt.Errorf("gorm: %d ID provided was invalid", id)
	}

	workflow := app.Workflow{ID: id}
	return c.DB.Delete(&workflow).Error
}

func (c *Client) AutoMigrate(object interface{}) error {
	if err := c.DB.AutoMigrate(object).Error; err != nil {
		return fmt.Errorf(err())
	}

	return nil
}
