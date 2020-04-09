// 🚀 Fiber is an Express inspired web framework written in Go with 💖
// 📌 API Documentation: https://fiber.wiki
// 📝 Github Repository: https://github.com/gofiber/fiber

package requestid

import (
	"unsafe"

	"github.com/gofiber/fiber"
	"github.com/google/uuid"
)

// Config defines the config for RequestID middleware
type Config struct {
	// Filter defines a function to skip middleware.
	// Optional. Default: nil
	Filter func(*fiber.Ctx) bool
	// Generator defines a function to generate an ID.
	// Optional. Default: func() string {
	//   return uuid.New().String()
	// }
	Generator func() string
}

// New adds an indentifier to the request
func New(config ...Config) func(*fiber.Ctx) {
	var cfg Config
	if len(config) > 0 {
		cfg = config[0]
	}
	// Set config default values
	if cfg.Generator == nil {
		cfg.Generator = func() string {
			return uuid.New().String()
		}
	}
	// Return middleware
	return func(c *fiber.Ctx) {
		// Get id from request
		rid := c.Get(fiber.HeaderXRequestID)
		// Create new id if empty
		if rid == "" {
			rid = cfg.Generator()
		}
		// Set new id to response
		c.Set(fiber.HeaderXRequestID, rid)
		// Bye
		c.Next()
	}
}

// Get returns the request identifier
func Get(c *fiber.Ctx) string {
	return getString(c.Fasthttp.Response.Header.Peek(fiber.HeaderXRequestID))
}

func getString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
