package main

const BOT_NAME_FA = "جرات - حقیقت"
const BOT_NAME_EN = "Truth - Dare"
const APIKEY = "5365644695:AAEzvWrPGxpki0H-3tNSrZ3AV-DiMxTMe5o"

// For cheating; that's a game for prank, but you can enter nothing here.
var GOOD_USERNAMES []string = []string{"0ne-zero"}

var MODES = []string{"RECRUITMENT", "CHOISE", "RESPONSE", "END"}
var CURRENT_MODE string

const START_FROM_BOT_SPECIFIER = "/start@HaghighatJorwat_bot"

// Formats
const WELCOME_MESSAGE = `سلام کاربر « %s » گرامی  به ربات جرات و حقیقت خوش آمدی 💐`
const MODERATOR_USERNAME_FORMAT = "🕹 کنترل کننده بازی :"
const PLAYERS_USERNAME_FORMAT = "👨‍👩‍👧 شرکت کنندگان :"
const YOU_ARE_NOT_MODERATOR = "تو کنترل کننده ی بازی نیستی, احمق \U0001F5FF \U0001F5FF \U0001F5FF"
const TURN_FORMAT = "نوبت @%s هست \U0001F5FF \U0001F5FF"
const TAKE_CHOISE = "سریع انتخاب کن گلابی \U0001F5FF \U0001F5FF"
const NO_LEADER = "کسی داره بازی رو کنترل نمیکنه کصخلا \U0001F5FF \U0001F5FF"
const JUST_MODERATOR_CAN_CONTROL = "فقط کنترل کننده میتونه بازی رو کنترل کنه.\n(ع جدی میگی کصخل...) \U0001F5FF \U0001F5FF \U0001F5FF"
const FAST_RESPONSE = "زر نزن سریع جواب بده بقیه منتظر هستن, گلابی \U0001F5FF \U0001F5FF"
const ITS_NOT_YOUR_TURN = "بیشعور نوبت تو نیست \U0001F5FF"
const YOU_ARE_MODERATOR = "تو کنترل کننده ی بازی هستی \U0001F5FF"

var IS_TRUTH_OR_DARE = ""

// Game information
var MODERATOR_USERNAME string
var PLAYERS_USERNAME []string
var WHO_IS_TURN string
var Truths []string
var Dares []string
