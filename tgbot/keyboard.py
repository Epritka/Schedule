from aiogram.types import ReplyKeyboardRemove, \
    ReplyKeyboardMarkup, KeyboardButton, \
    InlineKeyboardMarkup, InlineKeyboardButton

BTN_LOGIN = KeyboardButton("Авторизоваться")
BTN_CHANGE_GROUP = KeyboardButton("Сменить группу")
BTN_SHOW_SCHEDULE = KeyboardButton("Посмотреть расписание")


DEFAULT = ReplyKeyboardMarkup(
    resize_keyboard=True, one_time_keyboard=True).add(BTN_LOGIN).add(BTN_CHANGE_GROUP)


SCHEDULE = ReplyKeyboardMarkup(
    resize_keyboard=True, one_time_keyboard=True).add(BTN_SHOW_SCHEDULE)
