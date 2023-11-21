package handler

import (
	"coss/pkg/pg"
	"coss/pkg/response"
	"coss/service/access/internal/controller"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func CreateObject() fiber.Handler {
	var (
		ctrl = controller.NewAccessController(pg.Client)
	)
	return func(ctx *fiber.Ctx) error {
		var (
			bucketName, objectName string
		)
		bucketName = ctx.Params("bucketName")
		if bucketName == "" {
			return response.Resp400(ctx, nil, "bucketName cannot be empty")
		}
		objectName = ctx.Params("objectName")
		if objectName == "" {
			return response.Resp400(ctx, nil, "objectName cannot be empty")
		}
		fileHeader, err := ctx.FormFile("object")
		if err != nil {
			return response.Resp400(ctx, nil, "object cannot be empty")
		}
		err = ctrl.CreateObject(bucketName, objectName, fileHeader)
		if err != nil {
			logrus.Errorf("failed create object, err: %v", err)
			return response.Resp500(ctx, nil, "failed create object", err.Error())
		}
		return response.Resp200(ctx, nil, "upload file success")
	}
}
