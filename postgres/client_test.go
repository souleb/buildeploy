package postgres

import (
	"os"
	"testing"

	"github.com/rs/zerolog/log"
)

var testClient *Client

func TestMain(m *testing.M) {
	testClient = NewClient()
	err := testClient.Open()
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to db.")
	}

	os.Exit(m.Run())
}
