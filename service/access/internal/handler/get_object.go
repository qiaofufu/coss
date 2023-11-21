package handler

import (
	"coss/pkg/pg"
	"coss/pkg/response"
	"coss/pkg/tools"
	"coss/service/access/internal/controller"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"net/http"
)

func GetObject() fiber.Handler {
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
		// TODO: Range 头支持

		data, err := ctrl.GetObjectWithData(bucketName, objectName)
		if err != nil {
			logrus.Errorf("get object failed, err: %v", err)
			return response.Resp500(ctx, nil, "get object failed")
		}
		ctx.Set("Content-Type", tools.GetContentType(objectName))
		ctx.Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s-%s", bucketName, objectName))
		_, err = ctx.Write(data.Data)
		return ctx.SendStatus(http.StatusOK)
	}
}
