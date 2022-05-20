package main

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

var WELCOME_KEYBOARD = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonURL("بازی در گروه \U0001F5FF", "https://t.me/HaghighatJorwat_bot?startgroup=add"),
	),
)

var RECRUITMENT_KEYBOARD = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("من رهبرم \U0001F5FF", "i_am_leader"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("منم بازی \U0001F97A", "i_am_in"),
		tgbotapi.NewInlineKeyboardButtonData("شروع بازی \U0001F608", "start_game"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("پایان بازی \U0001F3C1", "finish"),
	),
)
var GAME_KEYBOARD = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("سوال تصادفی \U0001F3B2", "random_question"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("حقیقت \U0001F925", "truth"),
		tgbotapi.NewInlineKeyboardButtonData("شجاعت \U0001F628", "dare"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("بیا پایین \U0001F447", "come_down"),
		tgbotapi.NewInlineKeyboardButtonData("منو قبلی \U0001F519", "home"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("پایان بازی \U0001F3C1", "finish"),
	),
)

var RESPONSE_KEYBOARD = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("جواب دادم \U0001F595", "responded"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("نفر بعدی \U0001F448", "next_person"),
		tgbotapi.NewInlineKeyboardButtonData("سوال بعدی \U0001F914", "next_question"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("منو قبلی \U0001F519", "return_menu"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("پایان بازی \U0001F3C1", "finish"),
	),
)
