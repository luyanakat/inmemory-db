package cmd

import (
	"errors"
	"time"

	"github.com/goombaio/namegenerator"
	"github.com/luyanakat/inmemory-db/db"
	"github.com/luyanakat/inmemory-db/internal"
	"github.com/luyanakat/inmemory-db/model"
)

type Store struct {
	DataStore *db.DataStore
}

func NewStore(rows int) (*Store, error) {
	ds, err := db.NewDataStore(rows)
	if err != nil {
		return nil, err
	}
	return &Store{
		DataStore: ds,
	}, nil
}

func (s *Store) Free() {
	s.DataStore = nil
}

func (s *Store) AddList(rows int) error {
	seed := time.Now().UTC().UnixNano()
	nameGenerator := namegenerator.NewNameGenerator(seed)

	for i := 0; i < rows; i++ {
		name := nameGenerator.Generate()
		key := internal.GetID()
		err := s.DataStore.Add(key, &model.Student{
			ID:     key,
			Name:   name,
			Age:    internal.RandomInt(18, 25),
			School: "HUST",
		})
		if err != nil {
			return errors.New(err.Error())
		}
	}
	return nil
}

func (s *Store) AddPeople(st *model.Student) error {
	if st == nil {
		return errors.New("people must not be empty")
	}
	if st.Name == "" {
		return errors.New("name must not be empty")
	}
	if st.Age < 18 {
		return errors.New("age must > 18")
	}

	return s.DataStore.Add(st.ID, st)
}

func (s *Store) SearchPeople(ID string) (*model.Student, error) {
	return s.DataStore.Read(ID)
}

func (s *Store) DeleteById(ID string) error {
	return s.DataStore.Delete(ID)
}

func (s *Store) UpdateById(ID string, st *model.Student) error {
	return s.DataStore.Update(ID, st)
}

func (s *Store) PrintAsOrder(order internal.Order) {
	s.DataStore.ListAll(order)
}

func (s *Store) PrintByAge(age int) {
	s.DataStore.ListByAge(age)
}