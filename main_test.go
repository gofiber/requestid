// 🚀 Fiber is an Express inspired web framework written in Go with 💖
// 📌 API Documentation: https://fiber.wiki
// 📝 Github Repository: https://github.com/gofiber/fiber

package requestid

import (
	"net/http"
	"testing"

	"github.com/gofiber/fiber"
)

func TestSet(t *testing.T) {
	app := fiber.New()
	app.Use(New(Config{Generator: func() string {
		return "rid"
	}}))
	app.Get("/", func(ctx *fiber.Ctx) {
		ctx.Body("hello world")
	})

	req, _ := http.NewRequest("GET", "http://example.com/", nil)
	res, err := app.Test(req)
	if err != nil {
		t.Fatal(err.Error())
	}
	defer res.Body.Close()

	eq(t, res.Header.Get(fiber.HeaderXRequestID), "rid")
}

func TestGetNewIdInHandler(t *testing.T) {
	var ridInRouter string
	app := fiber.New()
	app.Use(New(Config{Generator: func() string {
		return "rid"
	}}))
	app.Get("/", func(ctx *fiber.Ctx) {
		ridInRouter = Get(ctx)
		ctx.Body("hello world")
	})

	req, _ := http.NewRequest("GET", "http://example.com/", nil)
	res, err := app.Test(req)
	if err != nil {
		t.Fatal(err.Error())
	}
	defer res.Body.Close()

	eq(t, res.Header.Get(fiber.HeaderXRequestID), "rid", "request id in header not equal to config")
	eq(t, ridInRouter, "rid", "request id in router not equal to config")
}

func TestGetOldIdInHandler(t *testing.T) {
	var ridInRouter string
	app := fiber.New()
	app.Use(New(Config{Generator: func() string {
		return "rid"
	}}))
	app.Get("/", func(ctx *fiber.Ctx) {
		ridInRouter = Get(ctx)
		ctx.Body("hello world")
	})

	req, _ := http.NewRequest("GET", "http://example.com/", nil)
	req.Header.Set(fiber.HeaderXRequestID, "old rid")
	res, err := app.Test(req)
	if err != nil {
		t.Fatal(err.Error())
	}
	defer res.Body.Close()

	eq(t, res.Header.Get(fiber.HeaderXRequestID), "old rid", "id in res header not equal to req")
	eq(t, ridInRouter, "old rid", "rid in router not equal to req header")
}

func eq(t *testing.T, actual, expected string, msgAndArgs ...string) bool {
	t.Helper()
	if expected != actual {
		t.Fatal(msgAndArgs)
	}
	return true

}
