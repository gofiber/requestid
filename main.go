// 🚀 Fiber is an Express inspired web framework written in Go with 💖
// 📌 API Documentation: https://fiber.wiki
// 📝 Github Repository: https://github.com/gofiber/fiber

package requestid

import (
	"github.com/gofiber/fiber"
	"github.com/google/uuid"
)

// Config defines the config for RequestID middleware
type Config struct {
	// Skip defines a function to skip middleware.
	// Optional. Default: nil
	Skip func(*fiber.Ctx) bool
	// Generator defines a function to generate an ID.
	// Optional. Default: func() string {
	//   return uuid.New().String()
	// }
	Generator func() string
}

// New adds an indentifier to the request using the `X-Request-ID` header
func New(config ...Config) func(*fiber.Ctx) {
	// Init config
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
	// Return middleware handler
	return func(c *fiber.Ctx) {
		// Skip middleware if Skip returns true
		if cfg.Skip != nil && cfg.Skip(c) {
			c.Next()
			return
		}
		// Get value from RequestID
		rid := c.Get(fiber.HeaderXRequestID)
		// Create new ID
		if rid == "" {
			rid = cfg.Generator()
		}
		c.Next()
		// Set X-Request-ID
		c.Set(fiber.HeaderXRequestID, rid)
	}
}