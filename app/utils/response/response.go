package response

import (
	"github.com/gofiber/fiber/v2"
)

type MetaResponse struct {
	Meta Meta `json:"meta"`
}

type APIResponse struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func RenderJSONWithOutData(c *fiber.Ctx, code int, message string) error {
	metaResponse := MetaResponse{}
	metaResponse.Meta.Code = code
	metaResponse.Meta.Message = message

	c.Status(metaResponse.Meta.Code).JSON(metaResponse)
	return nil
}

func RenderJSON(c *fiber.Ctx, code int, message string, data interface{}) error {
	apiResponse := APIResponse{}
	apiResponse.Meta.Code = code
	apiResponse.Meta.Message = message
	if data != nil {
		apiResponse.Data = data
	}

	c.Status(apiResponse.Meta.Code).JSON(apiResponse)
	return nil
}

func Unauthorized(c *fiber.Ctx, args ...string) error {
	var message string
	if len(args) == 0 {
		message = "Unauthorized please login."
	} else {
		message = args[0]
	}

	return RenderJSONWithOutData(c, fiber.StatusUnauthorized, message)
}

func BadRequest(c *fiber.Ctx, args ...string) error {
	var message string
	if len(args) == 0 {
		message = "Parameter is invalid."
	} else {
		message = args[0]
	}

	return RenderJSONWithOutData(c, fiber.StatusBadRequest, message)
}

func NotFound(c *fiber.Ctx, args ...string) error {
	var message string
	if len(args) == 0 {
		message = "The URI requested is invalid or the resource requested does not exist."
	} else {
		message = args[0]
	}

	return RenderJSONWithOutData(c, fiber.StatusNotFound, message)
}

func UnprocessableEntity(c *fiber.Ctx, args ...string) error {
	var message string
	if len(args) == 0 {
		message = "Unprocessable Entity please check your body."
	} else {
		message = args[0]
	}

	return RenderJSONWithOutData(c, fiber.StatusUnprocessableEntity, message)
}

func InternetServerError(c *fiber.Ctx, args ...string) error {
	var message string
	if len(args) == 0 {
		message = "Something went wrong on ETrackings end."
	} else {
		message = args[0]
	}

	return RenderJSONWithOutData(c, fiber.StatusInternalServerError, message)
}

func ServerError(c *fiber.Ctx, args ...string) error {
	var message string
	if len(args) == 0 {
		message = "Bad gateway on ETrackings end."
	} else {
		message = args[0]
	}

	return RenderJSONWithOutData(c, fiber.StatusBadGateway, message)
}

func Timeout(c *fiber.Ctx, args ...string) error {
	var message string
	if len(args) == 0 {
		message = "Request Timeout on ETrackings."
	} else {
		message = args[0]
	}

	return RenderJSONWithOutData(c, fiber.StatusRequestTimeout, message)
}
