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

func TestAddPlayerThatExists(t *testing.T) {
	team := Team{
		Name:    "Team1",
		Players: []string{"Player1"},
	}

	team.AddPlayer("Player1")

	require.Equal(t, 1, len(team.Players))
}

func TestAddPlayerToNil(t *testing.T) {
	var team *Team
	err, _ := team.AddPlayer("Player1")
	require.NotNil(t, err)
}

func TestRemovePlayer(t *testing.T) {
	team := Team{
		Name:    "Team1",
		Players: []string{"Player1"},
	}

	team.RemovePlayer("Player1")

	require.Equal(t, 0, len(team.Players))
}
