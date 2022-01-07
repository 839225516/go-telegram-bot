package telegram

import tb "gopkg.in/tucnak/telebot.v2"

const (
	START_CMD = "/start"
	PING_CMD  = "/ping"
)

var Cmds = []tb.Command{
	{
		Text:        START_CMD,
		Description: "开始",
	},
	{
		Text:        PING_CMD,
		Description: "存活检测",
	},
}