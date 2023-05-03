package learning_go

type Team struct {
	Name    string
	Players []string
}

func (t *Team) AddPlayer(name string) {
	t.Players = append(t.Players, name)
}

func (t *Team) RemovePlayer(name string) {
	for i, v := range t.Players {
		if v == name {
			t.Players = append(t.Players[:i], t.Players[i+1:]...)
			break
		}
	}
}
