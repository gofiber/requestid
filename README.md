## RequestID
RequestID adds an indentifier to the response using the `X-Request-ID` header

### Install
```
go get -u github.com/gofiber/fiber
go get -u github.com/gofiber/requestid
```

### Config
| Property | Type | Description | Default |
| :--- | :--- | :--- | :--- |
| Skip | `func(*fiber.Ctx) bool` | Defines a function to skip middleware | `nil` |
| Generator | ` func() string` | Generator defines a function to generate an ID.e | `func() string {   return uuid.New().String() }` |

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
