package main

import "github.com/dynastymasra/server"

/**
 * Created by Dynastymasra
 * Name     : Dimas Ragil T
 * Email    : dynastymasra@gmail.com
 * LinkedIn : http://www.linkedin.com/in/dynastymasra
 * Github   : https://github.com/dynastymasra
 * Mobile and Backend Developer
 */

func main() {
  host := "localhost:4000"
  server := server.NewPersonServer(host)
  server.Run()
}
