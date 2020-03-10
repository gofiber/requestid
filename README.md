### Install
```
go get -u github.com/gofiber/fiber
go get -u github.com/gofiber/requestid
```
### Example
```go
package main

import (
  "github.com/gofiber/fiber"
  "github.com/gofiber/requestid"
)

func main() {
  app := fiber.New()

  app.Use(requestid.New())

  app.Get("/", func(c *fiber.Ctx) {
    c.Send("Hello, World!")
  })

  app.Listen(3000)
}
```
### Test
```curl
curl -I http://localhost:3000
```
