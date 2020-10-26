package postgres

import (
	"context"
	"fmt"

	_ "github.com/jackc/pgx/v4"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/souleb/buildeploy/app"
)

var (
	ErrInvalidID = errors.New("sql: ID provided was invalid")
)

type Client struct {
	DB       *sqlx.DB
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
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=%s",
		c.Host, c.Port, c.User, c.Password, c.DBname, c.Timezone)
	db, err := sqlx.Connect("pgx", psqlInfo)
	if err != nil {
		return errors.Wrap(err, "sqlx DB open failed")
	}

	c.DB = db

	return nil
}

// ReadPipeline will get a pipeline by ID.
func (c *Client) ReadPipeline(ctx context.Context, id uint, pipeline *app.Pipeline) error {
	err := c.DB.GetContext(ctx, &pipeline, "SELECT * FROM pipeline WHERE id == $1", id)
	if err != nil {
		return errors.Wrap(err, "sql: ID provided was invalid")
	}

	return nil
}

// ReadWorkflow will get a workflow by ID.
func (c *Client) ReadWorkflow(ctx context.Context, id uint, workflow *app.Workflow) error {
	err := c.DB.GetContext(ctx, &workflow, "SELECT * FROM workflow WHERE id == $1", id)
	if err != nil {
		return errors.Wrap(err, "sql: ID provided was invalid")
	}

	return nil
}

// ReadJob will get a job by ID.
func (c *Client) ReadJob(ctx context.Context, id uint, job *app.Job) error {
	err := c.DB.GetContext(ctx, &job, "SELECT * FROM job WHERE id == $1", id)
	if err != nil {
		return errors.Wrap(err, "sql: ID provided was invalid")
	}

	return nil
}

// CreateWorkflow will create the provided workflow and backfill data
// like the ID, CreatedAt, and UpdatedAt fields.
func (c *Client) CreateWorkflow(workflow *app.Workflow) (int64, error) {
	stmt, err := c.DB.Prepare("INSERT INTO workflow(name) VALUES($1)")
	if err != nil {
		return 0, errors.Wrap(err, "sql: creation failed")
	}

	res, err := stmt.Exec(workflow.Name)
	if err != nil {
		return 0, errors.Wrap(err, "sql: creation failed")
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		return 0, errors.Wrap(err, "sql: creation failed")
	}

	workflow.ID = lastId

	rowCnt, err := res.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "sql: creation failed")
	}
	return rowCnt, nil
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
