package learning_go

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddPlayer(t *testing.T) {
	team := Team{Name: "Team1"}
	team.AddPlayer("Player1")
	require.Equal(t, 1, len(team.Players))
}

func TestRemovePlayer(t *testing.T) {
	team := Team{
		Name:    "Team1",
		Players: []string{"Player1"},
	}

	team.RemovePlayer("Player1")

	require.Equal(t, 0, len(team.Players))
}
