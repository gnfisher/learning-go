package learning_go

type Store interface {
	Get(id int)
}

type ConcreteStore struct{}

func (cs *ConcreteStore) Get(id int) {
}

func (cs *ConcreteStore) Delete(id int) {
}

func GetStore() Store {
	return &ConcreteStore{}
}

func experiment() {
}
