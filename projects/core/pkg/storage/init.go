package storage

import (
	"core/pkg/proto/pb"
	"errors"
)

type Todo = pb.Todo

var store []*Todo

func init() {
	store = []*Todo{}
}

func GetAll() []*Todo {
	return store
}

func GetById(id string) (*Todo, int, error) {
	for index, todo := range store {
		if todo.Id == id {
			return todo, index, nil
		}
	}

	return nil, 0, errors.New("todo doesn't exist")
}

func Add(item *Todo) {
	store = append(store, item)
}

func Update(item *Todo) (*Todo, error) {
	id := item.Id
	todo, index, err := GetById(id)
	if err != nil {
		return nil, err
	}

	store[index] = item

	return todo, nil
}

func Remove(item *Todo) error {
	_, index, err := GetById(item.Id)
	if err != nil {
		return err
	}

	if len(store) > 1 {
		store = append(store[:index], store[index+1:]...)
	} else {
		store = []*Todo{}
	}

	return nil
}

func Drop() {
	store = []*Todo{}
}
