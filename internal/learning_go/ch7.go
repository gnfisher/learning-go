package learning_go

import (
	"errors"
)

type Team struct {
	Name    string
	Players []string
}

type Class Team

type Game struct {
	Name string
}

func (t *Team) AddPlayer(name string) (err error, val bool) {
	if t == nil {
		err = errors.New("team is nil")
		return err, val
	}

	for _, v := range t.Players {
		if v == name {
			val = true
			return err, val
		}
	}
	t.Players = append(t.Players, name)
	val = true
	return err, val
}

func (t *Team) RemovePlayer(name string) {
	for i, v := range t.Players {
		if v == name {
			t.Players = append(t.Players[:i], t.Players[i+1:]...)
			break
		}
	}
}

func (t *Team) toClass() Class {
	return Class(*t)
}

func (t *Team) toGame() Game {
	return Game{Name: t.Name}
}
