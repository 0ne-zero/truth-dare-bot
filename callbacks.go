package main

import (
	"fmt"
	"math/rand"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func start_game_callback(update tgbotapi.Update, bot *tgbotapi.BotAPI, senter_username string) {
	if MODERATOR_USERNAME == "" {
		// Send pop-up to user
		callback := tgbotapi.NewCallback(update.CallbackQuery.ID, NO_LEADER)
		if _, err := bot.Request(callback); err != nil {
			fmt.Println(err)
		}
		return
	}
	if senter_username != MODERATOR_USERNAME {
		// Send pop-up to user
		callback := tgbotapi.NewCallback(update.CallbackQuery.ID, YOU_ARE_NOT_MODERATOR)
		if _, err := bot.Request(callback); err != nil {
			fmt.Println(err)
		}
		return
	}
	turnSelection(update, bot)
}
func i_am_in_callback(senter_username string, update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	// Moderator wants to join to players, but her/his already is part of it
	if senter_username == MODERATOR_USERNAME {
		// Send pop-up to user
		callback := tgbotapi.NewCallback(update.CallbackQuery.ID, "تو کنترل کننده هستی, خودت باید تو بازی باشی کصخل; خودت جزو بازیکن ها هستی")
		if _, err := bot.Request(callback); err != nil {
			fmt.Println(err)
		}
		return
	}
	// user is in players list, so tell her/him
	if stringIsExistsInSlice(senter_username, PLAYERS_USERNAME) {
		// Send pop-up to user
		callback := tgbotapi.NewCallback(update.CallbackQuery.ID, "تو جزو بازیکنان هستی, کصخل")
		if _, err := bot.Request(callback); err != nil {
			fmt.Println(err)
		}
		return
	}
	// There isn't any problem, so add the who wants to be parts of players list
	PLAYERS_USERNAME = append(PLAYERS_USERNAME, senter_username)

	// Create moderator info
	moderator_info := fmt.Sprintf("%s\n%s", MODERATOR_USERNAME_FORMAT, MODERATOR_USERNAME)
	// Create players info
	players_info := fmt.Sprintf("%s\n%s", PLAYERS_USERNAME_FORMAT, breakStringSliceInLines(PLAYERS_USERNAME))
	// Finally text
	text := fmt.Sprintf("%s\n%s\n%s\n\n\n%s", BOT_NAME_FA, moderator_info, players_info, JUST_MODERATOR_CAN_CONTROL)
	var chat_id int64
	if update.MyChatMember != nil {
		chat_id = update.MyChatMember.Chat.ID
	} else {
		chat_id = update.FromChat().ChatConfig().ChatID
	}
	edit_msg := tgbotapi.NewEditMessageTextAndMarkup(chat_id, update.CallbackQuery.Message.MessageID, text, RECRUITMENT_KEYBOARD)
	if _, err := bot.Send(edit_msg); err != nil {
		fmt.Println(err)
		// Send pop-up to user
		callback := tgbotapi.NewCallback(update.CallbackQuery.ID, "نشد, یک مشکلی پیش اومد")
		if _, err := bot.Request(callback); err != nil {
			fmt.Println(err)
		}
		return
	}
	// Send pop-up to user
	callback := tgbotapi.NewCallback(update.CallbackQuery.ID, "تو جزو بازیکن ها شدی")
	if _, err := bot.Request(callback); err != nil {
		fmt.Println(err)
	}
}
func next_person_callback(update tgbotapi.Update, bot *tgbotapi.BotAPI, senter_username string) {
	if senter_username != WHO_IS_TURN {
		// Send pop-up to user
		callback := tgbotapi.NewCallback(update.CallbackQuery.ID, ITS_NOT_YOUR_TURN)
		if _, err := bot.Request(callback); err != nil {
			fmt.Println(err)
		}
		return
	}
	turnSelection(update, bot)
}
func responded_callback(update tgbotapi.Update, bot *tgbotapi.BotAPI, senter_username string) {
	if senter_username != WHO_IS_TURN {
		// Send pop-up to user
		callback := tgbotapi.NewCallback(update.CallbackQuery.ID, ITS_NOT_YOUR_TURN)
		if _, err := bot.Request(callback); err != nil {
			fmt.Println(err)
		}
		return
	}
	turnSelection(update, bot)
}
func truth_callback(update tgbotapi.Update, bot *tgbotapi.BotAPI, senter_username string) {
	if senter_username != WHO_IS_TURN {
		// Send pop-up to user
		callback := tgbotapi.NewCallback(update.CallbackQuery.ID, ITS_NOT_YOUR_TURN)
		if _, err := bot.Request(callback); err != nil {
			fmt.Println(err)
		}
		return
	}
	// Question
	var q string = Truths[rand.Intn(len(Truths))]
	text := fmt.Sprintf("%s\n\n%s\n\n%s", fmt.Sprintf(TURN_FORMAT, WHO_IS_TURN), q, FAST_RESPONSE)
	edit_msg := tgbotapi.NewEditMessageTextAndMarkup(update.FromChat().ChatConfig().ChatID, update.CallbackQuery.Message.MessageID, text, RESPONSE_KEYBOARD)
	if _, err := bot.Send(edit_msg); err != nil {
		fmt.Println(err)
		// Send pop-up to user
		callback := tgbotapi.NewCallback(update.CallbackQuery.ID, "نشد, یک مشکلی پیش اومد")
		if _, err := bot.Request(callback); err != nil {
			fmt.Println(err)
		}
		return
	}
	IS_TRUTH_OR_DARE = "TRUTH"
}
func dare_callback(update tgbotapi.Update, bot *tgbotapi.BotAPI, senter_username string) {
	if senter_username != WHO_IS_TURN {
		// Send pop-up to user
		callback := tgbotapi.NewCallback(update.CallbackQuery.ID, ITS_NOT_YOUR_TURN)
		if _, err := bot.Request(callback); err != nil {
			fmt.Println(err)
		}
		return
	}
	// Question
	var q string = Dares[rand.Intn(len(Dares))]
	text := fmt.Sprintf("%s\n\n%s\n\n%s", fmt.Sprintf(TURN_FORMAT, WHO_IS_TURN), q, FAST_RESPONSE)
	edit_msg := tgbotapi.NewEditMessageTextAndMarkup(update.FromChat().ChatConfig().ChatID, update.CallbackQuery.Message.MessageID, text, RESPONSE_KEYBOARD)
	if _, err := bot.Send(edit_msg); err != nil {
		fmt.Println(err)
		// Send pop-up to user
		callback := tgbotapi.NewCallback(update.CallbackQuery.ID, "نشد, یک مشکلی پیش اومد")
		if _, err := bot.Request(callback); err != nil {
			fmt.Println(err)
		}
		return
	}

	IS_TRUTH_OR_DARE = "DARE"
}
func random_question_callback(update tgbotapi.Update, bot *tgbotapi.BotAPI, senter_username string) {
	if senter_username != WHO_IS_TURN {
		// Send pop-up to user
		callback := tgbotapi.NewCallback(update.CallbackQuery.ID, ITS_NOT_YOUR_TURN)
		if _, err := bot.Request(callback); err != nil {
			fmt.Println(err)
		}
		return
	}
	// Question
	var q string

	// Random number between 0 and 1
	// 0 = truth
	// 1 = dare
	if rand.Intn(1) == 0 {
		q = Truths[rand.Intn(len(Truths))]
		IS_TRUTH_OR_DARE = "TRUTH"
	} else {
		q = Dares[rand.Intn(len(Dares))]
		IS_TRUTH_OR_DARE = "DARE"
	}

	text := fmt.Sprintf("%s\n\n%s\n\n%s", fmt.Sprintf(TURN_FORMAT, WHO_IS_TURN), q, FAST_RESPONSE)
	edit_msg := tgbotapi.NewEditMessageTextAndMarkup(update.FromChat().ChatConfig().ChatID, update.CallbackQuery.Message.MessageID, text, RESPONSE_KEYBOARD)
	if _, err := bot.Send(edit_msg); err != nil {
		fmt.Println(err)
		// Send pop-up to user
		callback := tgbotapi.NewCallback(update.CallbackQuery.ID, "نشد, یک مشکلی پیش اومد")
		if _, err := bot.Request(callback); err != nil {
			fmt.Println(err)
		}
		return
	}
}
func i_am_leader_callback(update tgbotapi.Update, bot *tgbotapi.BotAPI, senter_username string) {
	// Moderator already exists
	if MODERATOR_USERNAME != "" {
		// Send pop-up to user
		if senter_username == MODERATOR_USERNAME {
			text := fmt.Sprintf("رهبر خودتی پادشاه/ملکه")
			callback := tgbotapi.NewCallback(update.CallbackQuery.ID, text)
			if _, err := bot.Request(callback); err != nil {
				fmt.Println(err)
			}
			return
		} else {
			text := fmt.Sprintf("رهبر(%s) از قبل انتخاب شده بهش احترام بزار بیشعور", MODERATOR_USERNAME)
			callback := tgbotapi.NewCallback(update.CallbackQuery.ID, text)
			if _, err := bot.Request(callback); err != nil {
				fmt.Println(err)
			}
			return
		}
	} else {
		// Set user as moderator
		MODERATOR_USERNAME = senter_username
		// Add user to the players list
		PLAYERS_USERNAME = append(PLAYERS_USERNAME, senter_username)

		// Create moderator info
		moderator_info := fmt.Sprintf("%s\n%s", MODERATOR_USERNAME_FORMAT, MODERATOR_USERNAME)

		// Create players info
		players_info := fmt.Sprintf("%s\n%s", PLAYERS_USERNAME_FORMAT, breakStringSliceInLines(PLAYERS_USERNAME))

		// Finally text
		text := fmt.Sprintf("%s\n%s\n%s\n\n\n%s", BOT_NAME_FA, moderator_info, players_info, JUST_MODERATOR_CAN_CONTROL)

		edit_msg := tgbotapi.NewEditMessageTextAndMarkup(update.FromChat().ChatConfig().ChatID, update.CallbackQuery.Message.MessageID, text, RECRUITMENT_KEYBOARD)
		// Edit previous bot message and, show moderator
		if _, err := bot.Send(edit_msg); err != nil {
			fmt.Println(err)
		}
		// Send text to moderator
		callback := tgbotapi.NewCallback(update.CallbackQuery.ID, YOU_ARE_MODERATOR)
		if _, err := bot.Request(callback); err != nil {
			fmt.Println(err)
		}
		return
	}
}
func recruitment_callback(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	// Create moderator info
	moderator_info := fmt.Sprintf("%s\n%s", MODERATOR_USERNAME_FORMAT, MODERATOR_USERNAME)

	// Create players info
	players_info := fmt.Sprintf("%s\n%s", PLAYERS_USERNAME_FORMAT, breakStringSliceInLines(PLAYERS_USERNAME))

	// Finally text
	text := fmt.Sprintf("%s\n%s\n%s\n\n\n%s", BOT_NAME_FA, moderator_info, players_info, JUST_MODERATOR_CAN_CONTROL)
	var chat_id int64
	if update.MyChatMember != nil {
		chat_id = update.MyChatMember.Chat.ID
	} else {
		chat_id = update.FromChat().ChatConfig().ChatID
	}
	msg := tgbotapi.NewMessage(chat_id, text)

	// Set appropiate inline keyboard
	msg.ReplyMarkup = RECRUITMENT_KEYBOARD

	// Send message
	if _, err := bot.Send(msg); err != nil {
		fmt.Println(err)
		// Send pop-up to user
		callback := tgbotapi.NewCallback(update.CallbackQuery.ID, "نشد, یک مشکلی پیش اومد")
		if _, err := bot.Request(callback); err != nil {
			fmt.Println(err)
		}
		return
	}
}
func next_question_callback(update tgbotapi.Update, bot *tgbotapi.BotAPI, senter_username string) {
	if senter_username != WHO_IS_TURN {
		// Send pop-up to user
		callback := tgbotapi.NewCallback(update.CallbackQuery.ID, ITS_NOT_YOUR_TURN)
		if _, err := bot.Request(callback); err != nil {
			fmt.Println(err)
		}
		return
	}
	var q string
	if IS_TRUTH_OR_DARE == "TRUTH" {
		q = Truths[rand.Intn(len(Truths))]
	} else if IS_TRUTH_OR_DARE == "DARE" {
		q = Dares[rand.Intn(len(Dares))]
	}
	text := fmt.Sprintf("%s\n\n%s\n\n%s", fmt.Sprintf(TURN_FORMAT, WHO_IS_TURN), q, FAST_RESPONSE)
	edit_msg := tgbotapi.NewEditMessageTextAndMarkup(update.FromChat().ChatConfig().ChatID, update.CallbackQuery.Message.MessageID, text, RESPONSE_KEYBOARD)
	if _, err := bot.Send(edit_msg); err != nil {
		fmt.Println(err)
		// Send pop-up to user
		callback := tgbotapi.NewCallback(update.CallbackQuery.ID, "نشد, یک مشکلی پیش اومد")
		if _, err := bot.Request(callback); err != nil {
			fmt.Println(err)
		}
		return
	}
}
func return_menu_callback(update tgbotapi.Update, bot *tgbotapi.BotAPI, senter_username string) {
	if senter_username != WHO_IS_TURN {
		// Send pop-up to user
		callback := tgbotapi.NewCallback(update.CallbackQuery.ID, ITS_NOT_YOUR_TURN)
		if _, err := bot.Request(callback); err != nil {
			fmt.Println(err)
		}
		return
	}

	text := fmt.Sprintf("%s\n%s", fmt.Sprintf(TURN_FORMAT, senter_username), TAKE_CHOISE)
	edit_msg := tgbotapi.NewEditMessageTextAndMarkup(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID, text, GAME_KEYBOARD)
	if _, err := bot.Send(edit_msg); err != nil {
		fmt.Println(err)
		// Send pop-up to user
		callback := tgbotapi.NewCallback(update.CallbackQuery.ID, "نشد, یک مشکلی پیش اومد")
		if _, err := bot.Request(callback); err != nil {
			fmt.Println(err)
		}
		return
	}
	CURRENT_MODE = "CHOISE"
}
func finish_callback(update tgbotapi.Update, bot *tgbotapi.BotAPI, senter_username string) {
	if senter_username == MODERATOR_USERNAME {
		del_request := tgbotapi.NewDeleteMessage(update.FromChat().ChatConfig().ChatID, update.CallbackQuery.Message.MessageID)
		if _, err := bot.Request(del_request); err != nil {
			fmt.Println(err)
			// Send pop-up to user
			callback := tgbotapi.NewCallback(update.CallbackQuery.ID, "نشد, یک مشکلی پیش اومد")
			if _, err := bot.Request(callback); err != nil {
				fmt.Println(err)
			}
			return
		}
		MODERATOR_USERNAME = ""
		PLAYERS_USERNAME = nil
		WHO_IS_TURN = ""
	} else {
		// Send pop-up to user
		callback := tgbotapi.NewCallback(update.CallbackQuery.ID, YOU_ARE_NOT_MODERATOR)
		if _, err := bot.Request(callback); err != nil {
			fmt.Println(err)
		}
	}

}
func come_down_callback(update tgbotapi.Update, bot *tgbotapi.BotAPI, senter_username string) {
	if senter_username != MODERATOR_USERNAME && senter_username != WHO_IS_TURN {
		if senter_username != WHO_IS_TURN {
			// Send pop-up to user
			callback := tgbotapi.NewCallback(update.CallbackQuery.ID, "نه کنترل کننده ای و نه نوبت تو هست, گمشو کنار")
			if _, err := bot.Request(callback); err != nil {
				fmt.Println(err)
			}
			return
		}
	}
	chat_id := update.FromChat().ChatConfig().ChatID

	// Copy message
	new_msg := tgbotapi.NewMessage(chat_id, update.CallbackQuery.Message.Text)
	if update.CallbackQuery.Message.ReplyMarkup != nil {
		new_msg.ReplyMarkup = update.CallbackQuery.Message.ReplyMarkup
	}
	// Send copy message
	if _, err := bot.Send(new_msg); err != nil {
		fmt.Println(err)
		// Send pop-up to user
		callback := tgbotapi.NewCallback(update.CallbackQuery.ID, "نشد, یک مشکلی پیش اومد")
		if _, err := bot.Request(callback); err != nil {
			fmt.Println(err)
		}
		return
	}
	// Delete message
	del_msg := tgbotapi.NewDeleteMessage(chat_id, update.CallbackQuery.Message.MessageID)
	if _, err := bot.Request(del_msg); err != nil {
		fmt.Println(err)
		// Send pop-up to user
		callback := tgbotapi.NewCallback(update.CallbackQuery.ID, "نشد, یک مشکلی پیش اومد")
		if _, err := bot.Request(callback); err != nil {
			fmt.Println(err)
		}
		return
	}
}

func home_callback(update tgbotapi.Update, bot *tgbotapi.BotAPI, senter_username string) {
	if senter_username != WHO_IS_TURN {
		if senter_username != WHO_IS_TURN {
			// Send pop-up to user
			callback := tgbotapi.NewCallback(update.CallbackQuery.ID, ITS_NOT_YOUR_TURN)
			if _, err := bot.Request(callback); err != nil {
				fmt.Println(err)
			}
			return
		}
	}

	// Create moderator info
	moderator_info := fmt.Sprintf("%s\n%s", MODERATOR_USERNAME_FORMAT, MODERATOR_USERNAME)

	// Create players info
	players_info := fmt.Sprintf("%s\n%s", PLAYERS_USERNAME_FORMAT, breakStringSliceInLines(PLAYERS_USERNAME))

	// Finally text
	text := fmt.Sprintf("%s\n%s\n%s\n\n\n%s", BOT_NAME_FA, moderator_info, players_info, JUST_MODERATOR_CAN_CONTROL)

	edit_msg := tgbotapi.NewEditMessageTextAndMarkup(update.FromChat().ChatConfig().ChatID, update.CallbackQuery.Message.MessageID, text, RECRUITMENT_KEYBOARD)
	if _, err := bot.Send(edit_msg); err != nil {
		fmt.Println(err)
		// Send pop-up to user
		callback := tgbotapi.NewCallback(update.CallbackQuery.ID, "نشد, یک مشکلی پیش اومد")
		if _, err := bot.Request(callback); err != nil {
			fmt.Println(err)
		}
		return
	}
}
