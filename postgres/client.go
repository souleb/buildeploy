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

type queryParams struct {
	query string
	id    int64
	value interface{}
}

type execParams struct {
	insertCmd string
	value     []interface{}
}

// ReadByID will retrieve a row by its id
func (c *Client) ReadByID(ctx context.Context, params *queryParams) error {
	err := c.DB.GetContext(ctx, params.value, params.query, params.id)
	if err != nil {
		return errors.Wrap(err, "PipelineService Client: ID provided was invalid")
	}
	return nil
}

// Create will create the provided object and backfill data
// like the ID, CreatedAt, and UpdatedAt fields.
func (c *Client) Create(ctx context.Context, params *execParams) (int64, error) {
	stmt, err := c.DB.Prepare(params.insertCmd)
	if err != nil {
		return 0, errors.Wrap(err, "PipelineService Client: creation failed")
	}

	res, err := stmt.Exec(params.value...)
	if err != nil {
		return 0, errors.Wrap(err, "sql: creation failed")
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return 0, errors.Wrap(err, "sql: creation failed")
	}

	return lastID, nil
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

	lastID, err := res.LastInsertId()
	if err != nil {
		return 0, errors.Wrap(err, "sql: creation failed")
	}

	workflow.ID = lastID

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
