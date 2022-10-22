from aiogram.types import InlineKeyboardButton, InlineKeyboardMarkup

BTN_LOGIN = InlineKeyboardButton("Авторизоваться", callback_data="login")
BTN_CHANGE_GROUP = InlineKeyboardButton(
    "Сменить группу", callback_data="change_group")

DEFAULT = InlineKeyboardMarkup().add(BTN_LOGIN, BTN_CHANGE_GROUP)

BTN_TODAY_SCHEDULE = InlineKeyboardButton("Сегодня", callback_data="today")
BTN_TOMORROW_SCHEDULE = InlineKeyboardButton("Завтра", callback_data="tomorrow")

SELECT_DAY = InlineKeyboardMarkup().add(BTN_TODAY_SCHEDULE).add(BTN_TOMORROW_SCHEDULE)