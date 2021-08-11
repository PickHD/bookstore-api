package utils

import "github.com/gofiber/fiber/v2"

type Response struct {
	C *fiber.Ctx
}

func NewResponse(ctx *fiber.Ctx) Response {
	return Response{C: ctx}
}

func (res Response) ResponseFormatter(code int, message string, err error, result map[string]interface{}) error {
	ctx := res.C

	if code < 400 {
		err := ctx.Status(code).JSON(&fiber.Map{
			"success": true,
			"message": message,
			"error":   nil,
			"data":    result,
		})

		if err != nil {
			return err
		}

		return nil
	}

	err = ctx.Status(code).JSON(&fiber.Map{
		"success": false,
		"message": message,
		"error":   err,
		"data":    result,
	})

	if err != nil {
		return err
	}

	return nil
}
