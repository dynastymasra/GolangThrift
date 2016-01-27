package main

import (
  "fmt"
  "git.apache.org/thrift.git/lib/go/thrift"
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

func main() {
  socket, err := thrift.NewTSocket("localhost:4000")
  if err != nil {
    panic(err)
  }

  transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	transport := transportFactory.GetTransport(socket)
  client := microservice.NewPersonServiceClientFactory(transport, protocolFactory)
  defer client.Transport.Close()

  if err := client.Transport.Open(); err != nil {
    panic(err)
  }

  p1 := NewPersonInit(1, "Dimas", "Ragil T", "dynastymasra@gmail.com", 20, true)
  p1, err = client.Create(p1)
  if err != nil {
    panic(err)
  }
  p2, err := client.Read(p1.ID)
  if err != nil {
    panic(err)
  }
  fmt.Println(p2.ID, p2.Firstname, *p2.Lastname, *p2.Email, p2.Age, p2.Active)

  p3 := NewPersonInit(2, "Ratna", "Siwi Y", "ratnasiwi.yunanta@gmail.com", 20, true)
  p3, err = client.Create(p3)
  if err != nil {
    panic(err)
  }
  p4, err := client.Read(p3.ID)
  if err != nil {
    panic(err)
  }
  fmt.Println(p4.ID, p4.Firstname, *p4.Lastname, *p4.Email, p4.Age, p4.Active)

  p5 := NewPersonInit(1, "Dimas", "Ragil T", "dynastymasra@gmail.com", 20, false)
  p5, err = client.Update(p5)
  if err != nil {
    panic(err)
  }
  p6, err := client.GetAll()
  for _, person := range p6 {
    fmt.Println(person.ID, person.Firstname, *person.Lastname, *person.Email, person.Age, person.Active)
  }

  p7 := client.Destroy(p5.ID)
  if p7 != nil {
    panic(err)
  }
  p8, err := client.GetAll()
  for _, person := range p8 {
    fmt.Println(person.ID, person.Firstname, *person.Lastname, *person.Email, person.Age, person.Active)
  }
}

func NewPersonInit(id int32, firstname, lastname, email string, age int16, active bool) *microservice.Person {
	return &microservice.Person {
    ID: id,
    Firstname: firstname,
    Lastname: &lastname,
    Email: &email,
    Age: age,
    Active: active,
	}
}
