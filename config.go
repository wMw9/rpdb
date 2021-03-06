package rpdb

import (
	"os"
)

var (
	amqpUrl             = os.Getenv("AMQP_URL") // amqp://guest:guest@localhost:5672/
	IGSessionID         = os.Getenv("IG_SESSIONID")
	IGQueryHash         = os.Getenv("IG_QUERYHASH")
	POSTGRES_PASSWORD   = os.Getenv("POSTGRES_PASSWORD")
	POSTGRES_PASSWORD_2 = os.Getenv("POSTGRES_PASSWORD_2")
	POSTGRES_USER       = os.Getenv("POSTGRES_USER")
	POSTGRES_DB         = os.Getenv("POSTGRES_DB")
	POSTGRES_HOST       = os.Getenv("POSTGRES_HOST")
	POSTGRES_PORT       = os.Getenv("POSTGRES_PORT")
	tgBotTkn            = os.Getenv("TG_BOT_TKN")
	tgChanErr           = os.Getenv("TG_CHAN_ERR")
	vkAccessTkn         = os.Getenv("VK_ACCESS_TKN")
	socketURL           = os.Getenv("SOCKET_URL")
)

const (
	vkAPIVersion     = "5.122"
	vkUsersGetURL    = "https://api.vk.com/method/users.get"
	vkUsersGetFields = "status"
	vkWallGetURL     = "https://api.vk.com/method/wall.get"
	vkWallGetFilter  = "owner"
)
