package tgbotapi

import (
	"fmt"
	"testing"
)

func TestBotInfo(t *testing.T) {
	bot.GetMe()
	if bot.Me == nil || bot.Me.ErrorCode != 0 {
		t.Fatalf("Err: GetMe(): %#v\n", bot.Me)
	}
	bot.GetWebhookInfo()
	if bot.WebHook == nil || bot.WebHook.ErrorCode != 0 {
		t.Fatalf("Err: GetWebhookInfo(): %#v\n", bot.WebHook)
	}
	fmt.Printf("webhook: %v\n", bot.WebHook.IsWebHook())
	fmt.Printf("bot id: %v\n", bot.Me.Result.Id)
	fmt.Printf("bot username: %v\n", bot.Me.Result.UserName)
	fmt.Printf("bot firstname: %v\n", bot.Me.Result.FirstName)
}

func TestGetUpdates(t *testing.T) {
	updates := bot.GetUpdates()
	if updates.OK != true {
		t.Fatalf("Err: GetUpdates(): %#v\n", updates)
	}

	for update := range updates.Result {
		fmt.Printf("%#v\n", update)
	}
}
