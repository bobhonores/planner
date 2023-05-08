//go:build e2e

package tests

import (
	"context"
	"testing"

	v1 "github.com/bobhonores/planner/api/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ActionTestSuite struct {
	suite.Suite
}

func (s *ActionTestSuite) TestAddAction() {
	s.T().Run("adds a new action successfully", func(t *testing.T) {
		client := GetClient()
		resp, err := client.AddAction(
			context.Background(),
			&v1.AddActionRequest{
				Name:        "do your homework",
				Description: "this activity should be made by 10 pm",
				Done:        false,
			},
		)
		assert.NoError(s.T(), err)
		assert.NotEmpty(s.T(), resp.Action.Id)
		assert.Equal(s.T(), "do your homework", resp.Action.Name)
		assert.Equal(s.T(), "this activity should be made by 10 pm", resp.Action.Description)
	})
}

func TestActionService(t *testing.T) {
	suite.Run(t, new(ActionTestSuite))
}
