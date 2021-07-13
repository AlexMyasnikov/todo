package data

import (
	pb "github.com/ChuvashPeople/todo/services"
)

type Todo struct {
	Id          int64
	Name        string
	Description string
	Status      bool
}

type FakeDb struct {
	Todos []Todo
}

func (db *FakeDb) Create(r *pb.CreateRequest) Todo {
	var id int64 = 1
	if len(db.Todos) > 0 {
		id = db.Todos[0].Id
		for _, t := range db.Todos {
			if t.Id > id {
				id = t.Id
			}
		}
		id++
	} else {
		id = 1
	}

	todo := Todo{Id: id, Name: r.Name, Description: r.Description, Status: false}
	db.Todos = append(db.Todos, Todo{Id: id, Name: r.Name, Description: r.Description, Status: false})
	return todo
}

func (db *FakeDb) Delete(id int64) {
	var exist = false
	for _, t := range db.Todos {
		if t.Id == id {
			exist = true
		}
	}
	if exist == false {
		panic("To do with this ID doesn't exist")
	}
	db.Todos = append(db.Todos[:id], db.Todos[id+1:]...)
}

func (db *FakeDb) Get(r *pb.GetByIdRequest) Todo {
	var todo Todo
	for _, t := range db.Todos {
		if r.Id == t.Id {
			todo = t
			break
		}
	}

	return todo

}
