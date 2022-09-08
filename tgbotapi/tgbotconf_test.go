package tgbotapi

var token string = "123456789:token"
var bot TeleBot = TeleBot{
	Api:   "https://api.telegram.org/bot%s/%s",
	Token: token,
}
