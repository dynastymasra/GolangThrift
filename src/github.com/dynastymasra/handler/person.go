package handler

import (
	"fmt"
	"log"
	"github.com/dynastymasra/microservice"
)

/**
 * Created by Dynastymasra
 * Name     : Dimas Ragil T
 * Email    : dynastymasra@gmail.com
 * LinkedIn : http://www.linkedin.com/in/dynastymasra
 * Github   : https://github.com/dynastymasra
 * Mobile and Backend Developer
 */

type PersonHandler struct {
	persons map[int32]*microservice.Person
}

func NewPersonHandler() *PersonHandler {
	return &PersonHandler {
		persons: make(map[int32]*microservice.Person),
	}
}

func (personHandler *PersonHandler) Create(person *microservice.Person) (*microservice.Person, error) {
	log.Println("Create new person...")
	personHandler.persons[person.ID] = person
	return person, nil
}

func (personHandler *PersonHandler) Read(id int32) (*microservice.Person, error) {
	log.Println("Read person...")
	person, err := personHandler.persons[id]
	if !err {
		log.Printf("Person with ID %s does not exist", id)
		return nil, fmt.Errorf("Person with ID %s does not exist", id)
	}
	return person, nil
}

func (personHandler *PersonHandler) Update(person *microservice.Person) (*microservice.Person, error) {
	log.Println("Update person...")
	personHandler.persons[person.ID] = person
	return person, nil
}

func (personHandler *PersonHandler) Destroy(id int32) error {
	log.Println("Delete person...")
	if _, ok := personHandler.persons[id]; ok {
		delete(personHandler.persons, id)
	}
	return nil
}

func (personHandler *PersonHandler) GetAll() ([]*microservice.Person, error) {
	log.Println("Get all person...")
	var persons []*microservice.Person
	for _, person := range personHandler.persons {
		persons = append(persons, person)
	}
	return persons, nil
}
