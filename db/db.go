package db

import (
	"errors"
	"fmt"
	"sort"

	"github.com/luyanakat/inmemory-db/internal"
	"github.com/luyanakat/inmemory-db/model"
)

type DataStore struct {
	store  map[string]*model.Student
	length int
}

func NewDataStore(rows int) (*DataStore, error) {
	if rows <= 0 {
		return nil, errors.New("rows must > 0")
	}
	return &DataStore{
		store:  make(map[string]*model.Student, rows),
		length: rows,
	}, nil
}

func (ds *DataStore) Add(key string, val *model.Student) error {
	if key == "" {
		return errors.New("key can't be empty")
	}
	if _, ok := ds.store[key]; ok {
		return fmt.Errorf("key %s is duplicated,", key)
	}
	if len(ds.store) == ds.length {
		return fmt.Errorf("the data store is full, only store %d rows", ds.length)
	}

	ds.store[key] = val
	return nil
}

func (ds *DataStore) Update(key string, val *model.Student) error {
	if _, ok := ds.store[key]; ok {
		ds.store[key] = val
	}
	if _, ok := ds.store[key]; !ok {
		return errors.New("key not exist")
	}
	return nil
}

func (ds *DataStore) Delete(key string) error {
	if _, ok := ds.store[key]; !ok {
		return errors.New("key not exist")
	}
	delete(ds.store, key)
	return nil
}

func (ds *DataStore) Read(key string) (*model.Student, error) {
	if _, ok := ds.store[key]; !ok {
		return nil, errors.New("key not exist")
	}
	return ds.store[key], nil
}

func (ds *DataStore) ListAll(order internal.Order) {
	keys := make([]string, 0)

	for key := range ds.store {
		keys = append(keys, key)
	}

	if order == internal.Asc {
		sort.Strings(keys)
	} else {
		sort.Sort(sort.Reverse(sort.StringSlice(keys)))
	}

	for _, key := range keys {
		fmt.Println(ds.store[key].String())
	}
}

func (ds *DataStore) ListByAge(age int) {
	keys := make([]string, 0)

	for key := range ds.store {
		keys = append(keys, key)
	}

	for _, key := range keys {
		if ds.store[key].Age == age {
			fmt.Println(ds.store[key].String())
		}
	}
}
