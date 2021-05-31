package controllers

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"

	"github.com/saharak-manoo/etracking-sdk-go/etracking"
	"github.com/saharak/etrackings-linebot/app/services"
	"github.com/saharak/etrackings-linebot/app/utils/linebot"
	"github.com/saharak/etrackings-linebot/app/utils/response"
)

func (server *Server) LineBotHooks(c *fiber.Ctx) error {
	lineBotService := services.LineBot{}
	events, _ := lineBotService.Events(c.Request())

	if len(events) > 0 {
		for _, event := range events {
			replyToken := event.ReplyToken

			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					etrackingApi, err := etracking.New(os.Getenv("ETRACKINGS_API_KEY"), os.Getenv("ETRACKINGS_KEY_SECRET"))
					if err != nil {
						fmt.Println("Errors ", err.Error())
					}

					resp, _ := etrackingApi.Find("jt-express", message.Text)
					if resp.Meta.Code == 200 {
						data := resp.Data

						lineBotService.ReplyStrMessage(replyToken, fmt.Sprintf("ติดตามพัสดุของ %s\nเลขพัสดุ: %s\nสถานะล่าสุด: %s", data.Courier, data.TrackingNo, data.CurrentStatus))
					} else {
						lineBotService.ReplyStrMessage(replyToken, "ไม่มีข้อมูลพัสดุ")
					}
				default:
					lineBotService.ReplyStrMessage(replyToken, "โปรดกรอกเลขพัสดุของคุณ")
				}
			}
		}
	}

	return response.RenderJSONWithOutData(c, fiber.StatusOK, "LineBot is OK")
}
