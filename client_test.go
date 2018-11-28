package libseat

import (
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

type libseatTestSuite struct {
	suite.Suite
	c *Client
}

func (s *libseatTestSuite) SetupSuite() {
	server := os.Getenv("SEAT_API_SERVER")
	apikey := os.Getenv("SEAT_API_KEY")
	if server == "" || apikey == "" {
		s.FailNow("Cannot run test suite without SEAT_API_SERVER and SEAT_API_KEY undefined")
	}

	client, err := NewClient(server, apikey)
	if err != nil {
		s.FailNow("Failed to create client", err)
	}
	s.c = client
}

func TestLibSeat(t *testing.T) {
	suite.Run(t, new(libseatTestSuite))
}
