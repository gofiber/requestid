# Request ID

![Release](https://img.shields.io/github/release/gofiber/requestid.svg)
[![Discord](https://img.shields.io/badge/discord-join%20channel-7289DA)](https://gofiber.io/discord)
![Test](https://github.com/gofiber/requestid/workflows/Test/badge.svg)
![Security](https://github.com/gofiber/requestid/workflows/Security/badge.svg)
![Linter](https://github.com/gofiber/requestid/workflows/Linter/badge.svg)

Adds an indentifier to the response using the `X-Request-ID` header

### Install
```
go get -u github.com/gofiber/fiber
go get -u github.com/gofiber/requestid
```

### Signature
```go
requestid.New(config ...requestid.Config) func(*fiber.Ctx)
```

### Config
| Property | Type | Description | Default |
| :--- | :--- | :--- | :--- |
| Skip | `func(*fiber.Ctx) bool` | A function to skip the middleware | `nil` |
| Generator | ` func() string` | A function to generate an ID.e | `return uuid.New().String()` |

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
    c.Send(requestid.Get(c))
  })

  app.Listen(3000)
}
```
### Test
```curl
curl -I http://localhost:3000
```
