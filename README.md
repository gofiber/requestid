## BasicAuth

Basic auth middleware provides an HTTP basic authentication. It calls the next handler for valid credentials and "401 - Unauthorized" for missing or invalid credentials.

**Signature**

```go
middleware.BasicAuth(config ...BasicAuthConfig) func(*Ctx)
```

**Config**

| Property | Type | Description | Default |
| :--- | :--- | :--- | :--- |
| Skip | `func(*Ctx) bool` | Defines a function to skip middleware | `nil` |
| Users | `map[string][string]` | Users defines the allowed credentials | `nil` |
| Realm | `string` | Realm is a string to define the realm attribute | `Restricted` |

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
