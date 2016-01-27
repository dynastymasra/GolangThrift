package server

import (
  "log"
  "git.apache.org/thrift.git/lib/go/thrift"
  "github.com/dynastymasra/microservice"
  "github.com/dynastymasra/handler"
)

/**
 * Created by Dynastymasra
 * Name     : Dimas Ragil T
 * Email    : dynastymasra@gmail.com
 * LinkedIn : http://www.linkedin.com/in/dynastymasra
 * Github   : https://github.com/dynastymasra
 * Mobile and Backend Developer
 */

type PersonServer struct {
  host string
  handler *handler.PersonHandler
  processor *microservice.PersonServiceProcessor
  transport *thrift.TServerSocket
  transportFactory thrift.TTransportFactory
  protocolFactory *thrift.TBinaryProtocolFactory
  server *thrift.TSimpleServer
}

func NewPersonServer(host string) *PersonServer {
  handler := handler.NewPersonHandler()
  processor := microservice.NewPersonServiceProcessor(handler)
  transport, err := thrift.NewTServerSocket(host)
  if err != nil {
    panic(err)
  }

  transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
  protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
  server := thrift.NewTSimpleServer4(processor, transport, transportFactory, protocolFactory)
  return &PersonServer {
    host: host,
    handler: handler,
    processor: processor,
    transport: transport,
    transportFactory: transportFactory,
    protocolFactory: protocolFactory,
    server: server,
  }
}

func (personServer *PersonServer) Run() {
  log.Printf("Server listening on %s...\n", personServer.host)
  personServer.server.Serve()
}

func (personServer *PersonServer) Stop() {
  log.Println("Stopping server...")
  personServer.server.Stop()
}
