//go:build e2e

package tests

import (
	"context"
	"testing"

	v1 "github.com/bobhonores/planner/api/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func (s *ActionTestSuite) TestGetAction() {
	s.T().Run("gets an invalid argument", func(t *testing.T) {
		client := GetClient()
		_, err := client.GetAction(
			context.Background(),
			&v1.GetActionRequest{
				Id: "invalid-id",
			},
		)
		assert.Error(s.T(), err)

		st := status.Convert(err)
		assert.Equal(s.T(), codes.InvalidArgument, st.Code())
	})

	s.T().Run("gets a non-found action", func(t *testing.T) {
		client := GetClient()
		_, err := client.GetAction(
			context.Background(),
			&v1.GetActionRequest{
				Id: "b304da2d-77b4-491c-a532-2ce17c0eceef",
			},
		)
		assert.Error(s.T(), err)

		st := status.Convert(err)
		assert.Equal(s.T(), codes.NotFound, st.Code())
	})
}

func TestActionService(t *testing.T) {
	suite.Run(t, new(ActionTestSuite))
}
