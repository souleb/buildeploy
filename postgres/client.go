package postgres

import (
	"context"
	"fmt"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

var (
	ErrInvalidID = errors.New("sql: ID provided was invalid")
)

type Client struct {
	db       *sqlx.DB
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
		return errors.Wrap(err, "sqlx db open failed")
	}

	c.db = db

	return nil
}

func (c *Client) Close() error {
	err := c.db.Close()
	if err != nil {
		return err
	}

	return nil
}

type queryParams struct {
	query string
	id    int64
	value interface{}
}

type execParams struct {
	insertCmd string
	values    []interface{}
}

// readByID will retrieve a row by its id
func (c *Client) readByID(ctx context.Context, params *queryParams) error {
	err := c.db.GetContext(ctx, params.value, params.query, params.id)
	if err != nil {
		return errors.Wrap(err, "PipelineService Client: ID provided was invalid")
	}
	return nil
}

// create will create the provided object and backfill data
// like the ID, CreatedAt, and UpdatedAt fields.
func (c *Client) create(ctx context.Context, params *execParams) (int64, error) {
	stmt, err := c.db.Preparex(params.insertCmd)
	if err != nil {
		return 0, errors.Wrap(err, "PipelineService Client: creation failed")
	}

	var id int64
	err = stmt.GetContext(ctx, &id, params.values...)

	if err != nil {
		return 0, errors.Wrap(err, "sql: error during creation")
	}

	return id, nil
}
