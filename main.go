package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	loadTruthsAndDares()

	bot, err := tgbotapi.NewBotAPI(APIKEY)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	// Loop through each update.
	for update := range updates {
		// If this field isn't nil, so that mean bot added in a group
		if update.MyChatMember != nil {
			recruitment_callback(update, bot)
			continue
		}
		// Check if we've gotten a message update.
		if update.Message != nil {
			if update.Message.IsCommand() {
				// If sent text equals to START_FROM_BOT_SPECIFIER means, someone start bot from group and without robot link.
				if update.Message.Text == START_FROM_BOT_SPECIFIER {
					start_handler(update, bot)
					continue
				}
				// Seperate real command and strange command like above command that handled
				switch update.Message.Text {
				case "/start":
					start_handler(update, bot)
					continue
				}
			}
		} else if update.CallbackQuery != nil {
			senter_username := update.CallbackQuery.From.UserName
			switch update.CallbackQuery.Data {
			// RECRUITMENT
			case "i_am_in":
				i_am_in_callback(senter_username, update, bot)
			case "start_game":
				start_game_callback(update, bot, senter_username)
			case "i_am_leader":
				i_am_leader_callback(update, bot, senter_username)
			// GAME
			case "random_question":
				random_question_callback(update, bot)
			case "truth":
				truth_callback(update, bot)
			case "dare":
				dare_callback(update, bot)
			case "come_down":
				come_down_callback(update, bot, senter_username)
			case "finish":
				finish_callback(update, bot, senter_username)
			case "home":
				home_callback(update, bot, senter_username)
			// RESPONSE
			case "responded":
				turnSelection(update, bot, senter_username)
			case "next_person":
				next_person_callback(update, bot, senter_username)
			case "next_question":
				next_question_callback(update, bot, senter_username)
			case "return_menu":
				return_menu_callback(update, bot, senter_username)
			}
		}
		continue
	}
}

func turnSelection(update tgbotapi.Update, bot *tgbotapi.BotAPI, senter_username string) {
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

	// Select a player for play the game
	randomSelectionPlayers()

	// Edit previous bot's message to a game keyboard
	text := fmt.Sprintf("%s\n%s", fmt.Sprintf(TURN_FORMAT, WHO_IS_TURN), TAKE_CHOISE)
	edit_msg := tgbotapi.NewEditMessageTextAndMarkup(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID, text, GAME_KEYBOARD)
	if _, err := bot.Send(edit_msg); err != nil {
		fmt.Println(err)
	}
	CURRENT_MODE = "CHOISE"
	return
}
func start_handler(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	// Bot started in a group
	if update.Message.Chat.Type == "group" || update.Message.Text == START_FROM_BOT_SPECIFIER {
		if CURRENT_MODE == "END" || CURRENT_MODE == "" {
			recruitment_callback(update, bot)
			return
		}
	} else if update.Message.Chat.Type == "private" {
		// This text will send
		text := fmt.Sprintf(WELCOME_MESSAGE, update.SentFrom().UserName)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
		// Set appropiate message keyboard
		msg.ReplyMarkup = WELCOME_KEYBOARD
		// Send it
		_, err := bot.Send(msg)
		if err != nil {
			fmt.Println(err)
		}
		return
	}
}

// NOT EXACTLY RANDOM - SEE VARS.GOOD_USERNAMES
func randomSelectionPlayers() {
	// if len(PLAYERS_USERNAME) < 2 {
	// 	// Send pop-up to user
	// 	callback := tgbotapi.NewCallback(update.CallbackQuery.ID, "برای شروع بازی حداقل دو بازیکن نیاز گلابی")
	// 	if _, err := bot.Request(callback); err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	return
	// }
	retry_time := 0

	result := PLAYERS_USERNAME[rand.Intn(len(PLAYERS_USERNAME))]

	for stringIsExistsInSlice(result, GOOD_USERNAMES) {
		if retry_time > 3 {
			WHO_IS_TURN = result
			return
		}

		result = PLAYERS_USERNAME[rand.Intn(len(PLAYERS_USERNAME))]
		retry_time += 1
	}
	WHO_IS_TURN = result
}
