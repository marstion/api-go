package tgbotapi

import (
	"encoding/json"
	"fmt"

	"github.com/marstion/api-go/request"
)

var header map[string][]string = map[string][]string{"Content-Type": {"application/json"}}

type TeleBot struct {
	// bot 类
	offset  int
	Api     string // "https://api.telegram.org/bot%s/%s"
	Token   string
	Me      *me
	WebHook *webHook
}

func (rec *TextReceiveResult) ToMsg() (msg *TextMessage) {
	if rec.CallbackQuery != nil {
		msg = &TextMessage{
			Text:        rec.CallbackQuery.Message.Text,
			ChatID:      rec.CallbackQuery.Message.Chat.ID,
			MessageID:   rec.CallbackQuery.Message.MessageID,
			ReplyMarkup: rec.CallbackQuery.Message.ReplyMarkup,
		}
	} else if rec.Message != nil {
		msg = &TextMessage{
			Text:        rec.Message.Text,
			ChatID:      rec.Message.Chat.ID,
			MessageID:   rec.Message.MessageID,
			ReplyMarkup: rec.Message.ReplyMarkup,
		}
	}
	return
}

type TextMessage struct {
	// 发送消息体
	Text              string `json:"text"`                     // Text of the message to be sent
	ParseMode         string `json:"parse_mode"`               // Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in your bot's message.
	ChatID            int64  `json:"chat_id"`                  // Unique identifier for the target chat or username of the target channel (in the format
	MessageID         int32  `json:"message_id"`               // If the message edit
	ReplyToMessageID  int32  `json:"reply_to_message_id"`      // If the message is a reply, ID of the original message
	DisNotification   bool   `json:"disable_notification"`     // Sends the message silently. iOS users will not receive a notification, Android users will receive a notification with no sound. Other apps coming soon.
	DisWebPagePreview bool   `json:"disable_web_page_preview"` // Disables link previews for links in this message
	// inlinekeyboardbutton
	ReplyMarkup *TextReceiveResultMessageReplyMarkup `json:"reply_markup"` // Additional interface options. A JSON-serialized object for a custom reply keyboard, instructions to hide keyboard or to force a reply from the user.
}

type TextReceive struct {
	OK     bool `json:"ok"`
	Result []*TextReceiveResult
}

type TextReceiveResult struct {
	// 接收消息体
	UpdateID      int `json:"update_id"`
	Message       *textReceiveResultMessage
	CallbackQuery *TextReceiveResultCallbackRuery `json:"callback_query"`
}

type TextReceiveResultCallbackRuery struct {
	ID           string `json:"id"`
	From         *textReceiveResultMessageFrom
	Message      *textReceiveResultMessage
	ChatInstance string `json:"chat_instance"`
	Data         string
}

type textReceiveResultMessage struct {
	MessageID      int32 `json:"message_id"`
	Date           int
	Text           string
	From           *textReceiveResultMessageFrom
	Chat           *textReceiveResultMessageChat
	ReplyToMessage *textReceiveResultMessage            `json:"reply_to_message"`
	ReplyMarkup    *TextReceiveResultMessageReplyMarkup `json:"reply_markup"`
}

type textReceiveResultMessageFrom struct {
	ID           int    `json:"id"`
	IsBot        bool   `json:"is_bot"`
	FirstName    string `json:"first_name"`
	UserName     string `json:"username"`
	LanguageCode string `json:"language_code"`
}

type textReceiveResultMessageChat struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
	Type  string `json:"type"`
}

type TextReceiveResultMessageReplyMarkup struct {
	InlineKeyboard [][]*TextReceiveResultMessageReplyMarkupInlineKeyboard `json:"inline_keyboard"`
}

type TextReceiveResultMessageReplyMarkupInlineKeyboard struct {
	// https://core.telegram.org/bots/api#inlinekeyboardmarkup
	Text                         string `json:"text"` // Label text on the button
	Pay                          bool   // Optional. Specify True, to send a Pay button. NOTE: This type of button must always be the first button in the first row and can only be used in invoice messages.
	Url                          string // HTTP or tg:// URL to be opened when the button is pressed. Links tg://user?id=<user_id> can be used to mention a user by their ID without using a username, if this is allowed by their privacy settings.
	CallbackData                 string `json:"callback_data"`                    // Data to be sent in a callback query to the bot when button is pressed, 1-64 bytes
	SwitchinlineQuery            string `json:"switch_inline_query"`              // If set, pressing the button will prompt the user to select one of their chats, open that chat and insert the bot's username and the specified inline query in the input field. May be empty, in which case just the bot's username will be inserted. Note: This offers an easy way for users to start using your bot in inline mode when they are currently in a private chat with it. Especially useful when combined with switch_pm… actions - in this case the user will be automatically returned to the chat they switched from, skipping the chat selection screen.
	SwitchinlineQueryCurrentChat string `json:"switch_inline_query_current_chat"` // If set, pressing the button will insert the bot's username and the specified inline query in the current chat's input field. May be empty, in which case only the bot's username will be inserted. This offers a quick way for the user to open your bot in inline mode in the same chat - good for selecting something from multiple options.

	// web_app WebAppInfo Description of the Web App that will be launched when the user presses the button. The Web App will be able to send an arbitrary message on behalf of the user using the method answerWebAppQuery. Available only in private chats between a user and the bot.
	// login_url LoginUrl An HTTPS URL used to automatically authorize the user. Can be used as a replacement for the Telegram Login Widget.
	// callback_game	CallbackGame	Optional. Description of the game that will be launched when the user presses the button. NOTE: This type of button must always be the first button in the first row.
}

type TextResponse struct {
	Ok     bool `json:"ok"`
	Result *textResponseResult

	ErrorCode   int    `json:"error_code"`  // only response 400
	Description string `json:"description"` // only response 400
}

type textResponseResult struct {
	From      *textResponseResultFrom
	Chat      *textResponseResultChat
	Text      string
	MessageID int32 `json:"message_id"`
}

type textResponseResultFrom struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	UserName  string
}

type textResponseResultChat struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	UserName  string
	LastName  string `json:"last_name"`
	Type      string `json:"type"`
}

type me struct {
	Ok          bool
	Result      *meResult
	ErrorCode   int `json:"error_code"`
	Description string
}

type meResult struct {
	Id        int
	FirstName string `json:"first_name"`
	UserName  string
}

type webHook struct {
	Ok          bool
	Result      *webHookResult
	ErrorCode   int `json:"error_code"`
	Description string
}

type webHookResult struct {
	Url                  string
	HasCustomCertificate bool   `json:"has_custom_certificate"`
	PendingUpdateCount   int    `json:"pending_update_count"`
	MaxConnections       int    `json:"max_connections"`
	IpAddress            string `json:"ip_address"`
}

func (wh webHook) IsWebHook() bool {
	if wh.Result.Url == "" {
		return false
	} else {
		return true
	}
}

func (t *TeleBot) GetMe() {
	var uri string = "getMe"

	rsp := request.Request("POST", fmt.Sprintf(t.Api, t.Token, uri), header, "")
	json.Unmarshal(rsp, &t.Me)
}

func (t *TeleBot) GetWebhookInfo() {
	var uri string = "getWebhookInfo"

	rsp := request.Request("POST", fmt.Sprintf(t.Api, t.Token, uri), header, "")
	json.Unmarshal(rsp, &t.WebHook)

}

func (t *TeleBot) GetUpdates() (rec *TextReceive) {
	uri := "getUpdates"
	msg := &map[string]int{"offset": t.offset}

	body_byte, _ := json.Marshal(msg)
	rsp := request.Request("POST", fmt.Sprintf(t.Api, t.Token, uri), header, string(body_byte))
	json.Unmarshal(rsp, &rec)

	// 更新updateid
	if len(rec.Result) > 0 {
		t.offset = rec.Result[len(rec.Result)-1].UpdateID + 1
	}
	return
}

func (t TeleBot) SendTextMessage(msg *TextMessage) (response *TextResponse) {
	uri := "sendMessage"

	body_byte, _ := json.Marshal(msg)
	rsp := request.Request("POST", fmt.Sprintf(t.Api, t.Token, uri), header, string(body_byte))
	json.Unmarshal(rsp, &response)
	return
}

func (t TeleBot) EditMessageText(msg *TextMessage) (response *TextResponse) {
	uri := "editMessageText"

	body_byte, _ := json.Marshal(msg)
	rsp := request.Request("POST", fmt.Sprintf(t.Api, t.Token, uri), header, string(body_byte))
	json.Unmarshal(rsp, &response)
	return
}

func (t TeleBot) DeleteMessage(msg *TextMessage) (response *TextResponse) {
	uri := "deleteMessage"

	body_byte, _ := json.Marshal(msg)
	rsp := request.Request("POST", fmt.Sprintf(t.Api, t.Token, uri), header, string(body_byte))
	json.Unmarshal(rsp, &response)
	return
}
