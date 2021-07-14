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
	db.Todos = append(db.Todos, todo)
	return todo
}

func (db *FakeDb) Update(r *pb.UpdateRequest) Todo {
	var todo Todo
	for i, t := range db.Todos {
		if r.Id == t.Id {
			if len(db.Todos)-i == 1 {
				db.Todos = append(db.Todos[:i])
			} else {
				db.Todos = append(db.Todos[:i], db.Todos[i+1:]...)
			}
			todo = Todo{Id: r.Id, Name: r.Name, Description: r.Description}
			db.Todos = append(db.Todos, todo)
			break
		}
	}
	return todo
}

func (db *FakeDb) Delete(r *pb.DeleteRequest) bool {
	var exist = false
	for i, t := range db.Todos {
		if t.Id == r.Id {
			exist = true
			if len(db.Todos)-i == 1 {
				db.Todos = append(db.Todos[:i])
			} else {
				db.Todos = append(db.Todos[:i], db.Todos[i+1:]...)
			}
		}
	}
	return exist
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

func (db *FakeDb) GetAll() []Todo {
	return db.Todos
}

func (db *FakeDb) Done(r *pb.MarkAsDoneRequest) Todo {
	var todo Todo
	for i := range db.Todos {
		if db.Todos[i].Id == r.Id {
			db.Todos[i].Status = true
			todo = db.Todos[i]
			break
		}
	}
	return todo
}
