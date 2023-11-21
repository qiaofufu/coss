package middlewares

import "github.com/gofiber/fiber/v2"

func ParseBucketAndObjectName() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		//var (
		//	bucketName, objectName string
		//)
		//bucketName = ctx.Params("bucketName")
		//objectName = ctx.Params("objectName")
		return ctx.Next()
	}
}
