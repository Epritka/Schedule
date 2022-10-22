from aiogram.dispatcher.filters.state import State, StatesGroup
from aiogram.contrib.fsm_storage.memory import MemoryStorage
from aiogram import Bot, Dispatcher, executor, types
from aiogram.dispatcher import FSMContext
from dotenv import load_dotenv
from datetime import datetime
import logging
import asyncio
import requests
import pytz
import json
import os

import keyboard
import inline_keyboard

DAYS = {
    1: "Понедельник",
    2: "Вторник",
    3: "Среда",
    4: "Четверг",
    5: "Пятница",
    6: "Суббота",
    7: "Воскресенье"
}

SERVER_URL = "http://127.0.0.1:2513"
users = {}

load_dotenv()
TG_TOKEN = os.getenv("TG_TOKEN")

logging.basicConfig(level=logging.INFO)
bot = Bot(token=TG_TOKEN)
storage = MemoryStorage()

dp = Dispatcher(bot, storage=storage)


class Form(StatesGroup):
    name = State()


@dp.message_handler(commands="help")
async def show_help_message(message: types.Message):
    await message.answer(text=f"Этот бот позволяет смотреть расписание вашего любимого университета♥")


@dp.message_handler(commands=["start"])
async def show_start_message(message: types.Message):
    await message.answer(f"Привет, {message.from_user.full_name}!",
                         reply_markup=inline_keyboard.DEFAULT)


@dp.callback_query_handler(text="login")
async def process_callback_login(callback_query: types.CallbackQuery):
    await bot.answer_callback_query(callback_query.id)
    await Form.name.set()
    await bot.send_message(
        callback_query.from_user.id,
        text="Введите группу"
    )


@dp.callback_query_handler(text="today")
async def process_callback_today(callback_query: types.CallbackQuery):
    await bot.answer_callback_query(callback_query.id)

    login = str(callback_query.from_user.full_name)
    moscow_time = datetime.now(pytz.timezone('Europe/Moscow'))
    json_data = {
        "groupId": users[login],
        "date": str(int(moscow_time.timestamp())),
    }

    response = requests.get(SERVER_URL+"/schedule/day", json=json_data)

    if response.status_code == 200:
        data = json.loads(response.text)["data"]
        text = f"{DAYS[data['number']]}\n"

        for l in data["lessons"]:
            text += f"{l['time']['start']}-{l['time']['end']}\n"

            if "subGroupNumber" in l:
                text += f"{l['subGroupNumber']}\n"
            text += f"{l['name']}\n"
            text += f"{l['teacher']}\n"
            text += f"{l['auditorium']}\n"
            text += f"{l['type']}\n\n"

        await bot.send_message(
            callback_query.from_user.id,
            text=text,
            reply_markup=keyboard.SCHEDULE,
        )
    else:
        await callback_query.message.reply("Извините...Что-то пошло не так.")


@dp.message_handler(content_types="text")
async def process_text_show_schedule(message: types.Message):
    if message.text == "Посмотреть расписание":
        await bot.send_message(
            message.from_user.id,
            text="Выберите на какой день хотите посмотреть расписание",
            reply_markup=inline_keyboard.SELECT_DAY,
        )


@dp.message_handler(state=Form.name)
async def process_text_group_name(message: types.Message, state: FSMContext):
    """Process group name"""

    await state.finish()

    login = str(message.from_user.full_name)
    response = requests.post(SERVER_URL+"/user/login", json={
        "login": login,
        "groupName": message.text
    })

    if response.status_code == 200:
        data = json.loads(response.text)
        global users
        users[login] = data["data"]["groupId"]
        await bot.send_message(
            message.from_user.id,
            text="Принято. Можете узнать расписание.",
            reply_markup=keyboard.SCHEDULE,
        )

    elif response.status_code == 404:
        await message.reply("Проверьте данные и попробуйте еще раз.")
    else:
        await message.reply("Извините...Что-то пошло не так.")


# @dp.message_handler()
# async def echo_message(msg: types.Message):
#     await bot.send_message(msg.from_user.id, msg.text)

if __name__ == "__main__":
    executor.start_polling(dp, skip_updates=True)
