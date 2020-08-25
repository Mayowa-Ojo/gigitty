package utils

import (
	"github.com/gofiber/fiber"
)

// NotFoundError -
func NotFoundError(c *fiber.Ctx) {
	// send 404 error
	err := fiber.NewError(404, "Resource not found.")

	c.Next(err)
}

// ErrorHandler -
func ErrorHandler(ctx *fiber.Ctx, err error) {
	code := fiber.StatusInternalServerError
	r := NewResponse()

	if e, ok := err.(*fiber.Error); ok {
		r.JSONResponse(ctx, false, e.Code, e.Message, make(EmptyMap))
	} else {
		r.JSONResponse(ctx, false, code, "[Error]: Internal server error", make(EmptyMap))
	}
}
