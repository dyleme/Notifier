package telegram

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"

	"github.com/dyleme/Notifier/internal/service"
	"github.com/dyleme/Notifier/pkg/log"
)

const (
	defaultListLimit = 100
)

var defaultListParams = service.ListParams{
	Offset: 0,
	Limit:  defaultListLimit,
}

func onSelectErrorHandling(
	f func(ctx context.Context, b *bot.Bot, relatedMsgID int, chatID int64) error,
) func(ctx context.Context, b *bot.Bot, msg models.MaybeInaccessibleMessage, _ []byte) {
	return func(ctx context.Context, b *bot.Bot, msg models.MaybeInaccessibleMessage, _ []byte) {
		err := f(ctx, b, msg.Message.ID, msg.Message.Chat.ID)
		if err != nil {
			handleError(ctx, b, msg.Message.Chat.ID, err)
		}
	}
}

func errorHandling(f func(ctx context.Context, b *bot.Bot, msg *models.Message, bts []byte) error) func(ctx context.Context, b *bot.Bot, msg models.MaybeInaccessibleMessage, _ []byte) {
	return func(ctx context.Context, b *bot.Bot, msg models.MaybeInaccessibleMessage, bts []byte) {
		err := f(ctx, b, msg.Message, bts)
		if err != nil {
			handleError(ctx, b, msg.Message.Chat.ID, err)
		}
	}
}

func handleError(ctx context.Context, b *bot.Bot, chatID int64, err error) {
	log.Ctx(ctx).Error("error occurred", log.Err(err))
	if chatID == 0 {
		return
	}

	_, err = b.SendMessage(ctx, &bot.SendMessageParams{ //nolint:exhaustruct //no need to specify
		ChatID: chatID,
		Text:   "Server error occurred\n" + err.Error(),
	})
	if err != nil {
		log.Ctx(ctx).Error("cannot send error message", log.Err(err))
	}
}
